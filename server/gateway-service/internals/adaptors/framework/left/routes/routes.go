package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/ports"
)

type Adaptor struct {
	authApi ports.AuthApiPort
}

func Initialize(aApi ports.AuthApiPort) *Adaptor {
	return &Adaptor{
		authApi: aApi,
	}
}

func (ra *Adaptor) Routes() http.Handler {
	mux := chi.NewRouter()
	// Specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	logger := httplog.NewLogger("gochi.log", httplog.Options{
		JSON: true,
	})
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Use(httplog.RequestLogger(logger))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(2500 * time.Millisecond))
	mux.Use(middleware.Throttle(10))
	mux.Use(httprate.LimitByIP(100, 1*time.Minute))
	mux.Route("/api/v1/auth", ra.AuthRoutes)
	return mux
}
