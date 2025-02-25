package handlers

import (
	"encoding/json"
	"net/http"
	"stockapi/internal/application/dto"
	"stockapi/internal/application/services"
	"stockapi/internal/domain/stock"

	"github.com/gorilla/mux"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type StockHandler struct {
	stockService *services.StockService
}

func NewStockHandler(service *services.StockService) *StockHandler {
	return &StockHandler{
		stockService: service,
	}
}

func (h *StockHandler) HandleStocks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getStocks(w, r)
		case http.MethodPost:
			h.syncStocks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *StockHandler) HandleStockDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getStockDetail(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *StockHandler) getStocks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	recommended := r.URL.Query().Get("recommended")
	var stocks []*stock.Stock
	var err error

	if recommended == "true" {
		stocks, err = h.stockService.GetRecommendedStocks(ctx)
	} else {
		stocks, err = h.stockService.GetAllStocks(ctx)
	}

	if err != nil {
		http.Error(w, "Error fetching stocks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir entidades a DTOs
	stockResponses := make([]dto.StockResponse, len(stocks))
	for i, s := range stocks {
		stockResponses[i] = dto.ToStockResponse(s)
	}

	w.Header().Set(ContentType, ApplicationJSON)
	json.NewEncoder(w).Encode(stockResponses)
}

func (h *StockHandler) syncStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ApplicationJSON)

	ctx := r.Context()
	if err := h.stockService.SyncStocksFromAPI(ctx); err != nil {
		response := map[string]string{"error": "Error syncing stocks: " + err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]string{"message": "Stocks synchronized successfully"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *StockHandler) getStockDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	if symbol == "" {
		http.Error(w, "Symbol is required", http.StatusBadRequest)
		return
	}

	stock, err := h.stockService.GetStockBySymbol(r.Context(), symbol)
	if err != nil {
		http.Error(w, "Error fetching stock detail: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert entity to DTO
	stockResponse := dto.ToStockResponse(stock)

	w.Header().Set(ContentType, ApplicationJSON)
	json.NewEncoder(w).Encode(stockResponse)
}
