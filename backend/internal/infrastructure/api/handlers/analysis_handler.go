package handlers

import (
	"encoding/json"
	"net/http"
	"stockapi/internal/application/services"
)

type AnalysisHandler struct {
	analysisService *services.AnalysisApplicationService
}

func NewAnalysisHandler(service *services.AnalysisApplicationService) *AnalysisHandler {
	return &AnalysisHandler{
		analysisService: service,
	}
}

func (h *AnalysisHandler) HandleAnalysis() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()
		analyses, err := h.analysisService.AnalyzeAllStocks(ctx)
		if err != nil {
			http.Error(w, "Error analyzing stocks: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(analyses)
	}
}
