package api

import (
	"fmt"
	"log"
	"net/http"
	"stockapi/internal/application"
	"stockapi/internal/infrastructure/api/handlers"
	"stockapi/internal/infrastructure/api/middleware"
	"stockapi/internal/infrastructure/config"
)

type Server struct {
	config       *config.Config
	stockHandler *handlers.StockHandler
	router       *http.ServeMux
}

func NewServer(cfg *config.Config, app *application.StockApplication) *Server {
	server := &Server{
		config:       cfg,
		stockHandler: handlers.NewStockHandler(app.StockService),
		router:       http.NewServeMux(),
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Apply middleware to all routes
	s.router.Handle("/api/stocks", middleware.Chain(
		s.stockHandler.HandleStocks(),
		middleware.Logging,
		middleware.CORS,
	))

	s.router.Handle("/api/stocks/recommended", middleware.Chain(
		s.stockHandler.HandleStocks(),
		middleware.Logging,
		middleware.CORS,
	))
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Server starting on port %s", s.config.Port)
	return http.ListenAndServe(addr, s.router)
}
