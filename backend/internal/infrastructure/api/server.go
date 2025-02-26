package api

import (
	"fmt"
	"log"
	"net/http"
	"stockapi/internal/application"
	"stockapi/internal/infrastructure/api/handlers"
	"stockapi/internal/infrastructure/api/middleware"
	"stockapi/internal/infrastructure/config"

	"github.com/gorilla/mux"
)

type Server struct {
	config          *config.Config
	app             *application.StockApplication
	stockHandler    *handlers.StockHandler
	analysisHandler *handlers.AnalysisHandler
	router          *mux.Router
}

func NewServer(cfg *config.Config, app *application.StockApplication) *Server {
	server := &Server{
		config: cfg,
		app:    app,
		router: mux.NewRouter(),
	}

	if app != nil {
		server.stockHandler = handlers.NewStockHandler(app.StockService)
		server.analysisHandler = handlers.NewAnalysisHandler(app.AnalysisService)
	}

	server.setupRoutes()
	return server
}

func (s *Server) UpdateApplication(app *application.StockApplication) {
	s.app = app
	s.stockHandler = handlers.NewStockHandler(app.StockService)
	s.analysisHandler = handlers.NewAnalysisHandler(app.AnalysisService)
	s.setupRoutes()
}

func (s *Server) setupRoutes() {
	// API routes
	s.router.HandleFunc("/api/stocks", s.stockHandler.HandleStocks()).
		Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	s.router.HandleFunc("/api/stocks/recommended", s.analysisHandler.HandleAnalysis()).
		Methods(http.MethodGet, http.MethodOptions)

	s.router.HandleFunc("/api/stocks/{symbol}", s.stockHandler.HandleStockDetail()).
		Methods(http.MethodGet, http.MethodOptions)

	// Apply global middleware
	s.router.Use(middleware.Logging)
	s.router.Use(middleware.CORS(s.config))
	s.router.Use(middleware.RateLimit)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Server starting on port %s", s.config.Port)
	return http.ListenAndServe(addr, s.router)
}
