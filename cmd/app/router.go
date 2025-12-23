package main

import (
	"net/http"
	"time"

	httpin_integration "github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi/v5"
	config "github.com/jorgeAM/grpc-kata-order-service/cfg"
	orderHandler "github.com/jorgeAM/grpc-kata-order-service/internal/order/infrastructure/http"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/http/handler"
	"github.com/jorgeAM/grpc-kata-order-service/pkg/http/middleware"
)

func buildRouter(cfg *config.Config, deps *config.Dependencies) http.Handler {
	router := chi.NewRouter()

	httpin_integration.UseGochiURLParam("path", chi.URLParam)

	router.Use(
		middleware.RequestID,
		middleware.Logger(middleware.WithIgnoreRoutes("/health")),
		middleware.Recover,
		middleware.RealIP,
		middleware.CORS(middleware.DefaultCORSOptions),
		middleware.ResponseHeader("Content-Type", "application/json"),
		middleware.ResponseHeader("Accept", "application/json"),
		middleware.Timeout(15*time.Second),
	)

	router.Get("/health", handler.HealthCheck)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/orders", func(r chi.Router) {
			r.Post("/", orderHandler.CreateOrder(cfg, deps))
		})
	})

	return router
}
