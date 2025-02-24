package stockapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"stockapi/internal/domain/shared"
	"stockapi/internal/domain/stock"
	"strconv"
	"strings"
	"time"
)

type StockAPIClient struct {
	baseURL    string
	httpClient *http.Client
	authToken  string
	logger     shared.Logger
}

type apiResponse struct {
	Items    []stockDTO `json:"items"`
	NextPage string     `json:"next_page"`
}

type stockDTO struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

func NewStockAPIClient(baseURL, authToken string, logger shared.Logger) stock.StockAPIPort {
	return &StockAPIClient{
		baseURL:    baseURL,
		authToken:  authToken,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		logger:     logger,
	}
}

func (c *StockAPIClient) FetchStocks(ctx context.Context) ([]*stock.Stock, error) {
	var allStocks []*stock.Stock
	nextPage := "" // Initially empty

	for hasMorePages := true; hasMorePages; {
		// Build the URL
		url := c.baseURL + "/list"
		if nextPage != "" {
			url = fmt.Sprintf("%s/list?next_page=%s", c.baseURL, nextPage)
		}

		c.logger.Debug(ctx, "Starting request to the API", map[string]interface{}{
			"url": url,
		})

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			c.logger.Error(ctx, "Error creating request", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+c.authToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error fetching stocks: %w", err)
		}
		defer resp.Body.Close()

		var apiResp apiResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}

		stocks, err := c.convertToStocks(apiResp.Items)
		if err != nil {
			return nil, err
		}

		allStocks = append(allStocks, stocks...)

		c.logger.Info(ctx, "Successfully obtained stocks page", map[string]interface{}{
			"count_in_page": len(stocks),
			"total_so_far":  len(allStocks),
			"next_page":     apiResp.NextPage,
		})

		// Update the loop condition
		hasMorePages = apiResp.NextPage != ""
		nextPage = apiResp.NextPage
	}

	return allStocks, nil
}

func (c *StockAPIClient) convertToStocks(items []stockDTO) ([]*stock.Stock, error) {
	var stocks []*stock.Stock
	for _, item := range items {
		targetFrom, err := parseMoneyString(item.TargetFrom)
		if err != nil {
			return nil, fmt.Errorf("error parsing target from: %w", err)
		}

		targetTo, err := parseMoneyString(item.TargetTo)
		if err != nil {
			return nil, fmt.Errorf("error parsing target to: %w", err)
		}

		stockEntity, err := stock.NewStock(item.Ticker, targetFrom, targetTo)
		if err != nil {
			return nil, fmt.Errorf("error creating stock entity: %w", err)
		}

		stockEntity.Company = item.Company
		stockEntity.Action = item.Action
		stockEntity.Brokerage = item.Brokerage
		stockEntity.Rating.From = stock.Rating(item.RatingFrom)
		stockEntity.Rating.To = stock.Rating(item.RatingTo)

		// Parse time
		t, err := time.Parse(time.RFC3339, item.Time)
		if err != nil {
			return nil, fmt.Errorf("error parsing time: %w", err)
		}
		stockEntity.Time = t

		stocks = append(stocks, stockEntity)
	}
	return stocks, nil
}

func parseMoneyString(value string) (stock.Money, error) {
	// Remove the $ symbol if it exists
	value = strings.TrimPrefix(value, "$")

	// Remove the commas from the thousands
	value = strings.ReplaceAll(value, ",", "")

	var amount float64
	var currency string

	_, err := fmt.Sscanf(value, "%f %s", &amount, &currency)
	if err != nil {
		// If no currency is specified, we assume USD
		amount, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return stock.Money{}, fmt.Errorf("error parsing money string '%s': %w", value, err)
		}
		currency = "USD"
	}

	return stock.Money{
		Amount:   amount,
		Currency: currency,
	}, nil
}
