package handlers

import (
	"encoding/json"
	"net/http"
	"stockapi/internal/application/services"
	"stockapi/internal/domain/stock"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}

func (h *StockHandler) syncStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
