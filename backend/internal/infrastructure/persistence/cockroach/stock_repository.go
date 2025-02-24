package cockroach

import (
	"context"
	"fmt"
	"stockapi/internal/domain/shared"
	"stockapi/internal/domain/stock"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockRepository struct {
	db     *pgxpool.Pool
	logger shared.Logger
}

func NewStockRepository(ctx context.Context, dbURL string, logger shared.Logger) (stock.Repository, error) {
	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &StockRepository{
		db:     dbPool,
		logger: logger,
	}, nil
}

func (r *StockRepository) Save(ctx context.Context, stock *stock.Stock) error {
	// First we check if the stock already exists
	existingStock, err := r.FindByTicker(ctx, stock.Ticker)
	if err == nil && existingStock != nil {
		// The stock exists, we update instead of inserting
		r.logger.Debug(ctx, "Stock already exists, updating", map[string]interface{}{
			"ticker": stock.Ticker,
			"id":     stock.ID,
		})
        
        return r.Update(ctx, stock)
    }

	// If it doesn't exist or there was an error searching (not found), we proceed with the insertion
	r.logger.Debug(ctx, "Inserting new stock", map[string]interface{}{
		"ticker": stock.Ticker,
		"id":     stock.ID,
	})

    query := `
        INSERT INTO stocks (
            id, ticker, target_from_amount, target_from_currency,
            target_to_amount, target_to_currency, company,
            action, brokerage, rating_from, rating_to, time
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `

    _, err = r.db.Exec(ctx, query,
        stock.ID,
        stock.Ticker,
        stock.Target.From.Amount,
        stock.Target.From.Currency,
        stock.Target.To.Amount,
        stock.Target.To.Currency,
        stock.Company,
        stock.Action,
        stock.Brokerage,
        stock.Rating.From,
        stock.Rating.To,
        stock.Time,
    )

    if err != nil {
        return fmt.Errorf("error saving stock: %w", err)
    }

	r.logger.Info(ctx, "Stock saved successfully", map[string]interface{}{
		"ticker": stock.Ticker,
		"id":     stock.ID,
	})
    return nil
}

func (r *StockRepository) FindAll(ctx context.Context) ([]*stock.Stock, error) {
	query := `
        SELECT id, ticker, target_from_amount, target_from_currency,
               target_to_amount, target_to_currency, company,
               action, brokerage, rating_from, rating_to, time
        FROM stocks
        ORDER BY time DESC
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying stocks: %w", err)
	}
	defer rows.Close()

	var stocks []*stock.Stock
	for rows.Next() {
		var s stock.Stock
		err := rows.Scan(
			&s.ID,
			&s.Ticker,
			&s.Target.From.Amount,
			&s.Target.From.Currency,
			&s.Target.To.Amount,
			&s.Target.To.Currency,
			&s.Company,
			&s.Action,
			&s.Brokerage,
			&s.Rating.From,
			&s.Rating.To,
			&s.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning stock: %w", err)
		}
		stocks = append(stocks, &s)
	}

	return stocks, nil
}

func (r *StockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM stocks WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting stock: %w", err)
	}

	return nil
}

func (r *StockRepository) FindByTicker(ctx context.Context, ticker string) (*stock.Stock, error) {
	query := `
        SELECT id, ticker, target_from_amount, target_from_currency,
               target_to_amount, target_to_currency, company,
               action, brokerage, rating_from, rating_to, time
        FROM stocks
        WHERE ticker = $1
        LIMIT 1
    `

	var s stock.Stock
	err := r.db.QueryRow(ctx, query, ticker).Scan(
		&s.ID,
		&s.Ticker,
		&s.Target.From.Amount,
		&s.Target.From.Currency,
		&s.Target.To.Amount,
		&s.Target.To.Currency,
		&s.Company,
		&s.Action,
		&s.Brokerage,
		&s.Rating.From,
		&s.Rating.To,
		&s.Time,
	)

	if err != nil {
		return nil, fmt.Errorf("error finding stock by ticker: %w", err)
	}

	return &s, nil
}

func (r *StockRepository) Update(ctx context.Context, stock *stock.Stock) error {
	query := `
        UPDATE stocks SET
            target_from_amount = $1,
            target_from_currency = $2,
            target_to_amount = $3,
            target_to_currency = $4,
            company = $5,
            action = $6,
            brokerage = $7,
            rating_from = $8,
            rating_to = $9,
            time = $10
        WHERE id = $11
    `

	_, err := r.db.Exec(ctx, query,
		stock.Target.From.Amount,
		stock.Target.From.Currency,
		stock.Target.To.Amount,
		stock.Target.To.Currency,
		stock.Company,
		stock.Action,
		stock.Brokerage,
		stock.Rating.From,
		stock.Rating.To,
		stock.Time,
		stock.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating stock: %w", err)
	}

	return nil
}

func (r *StockRepository) Close(ctx context.Context) error {
	r.db.Close()
	return nil
}
