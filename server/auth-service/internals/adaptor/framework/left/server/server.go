package server

import (
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"
)

type Adaptor struct {
	config ports.ConfigPort
	routes ports.RoutesPort
}

func Initialize(cfg ports.ConfigPort, routes ports.RoutesPort) *Adaptor {
	return &Adaptor{
		config: cfg,
		routes: routes,
	}
}

func (server *Adaptor) Server() error {
	cfg := server.config.GetConfig()
	cfg.Logger.Printf("Server is running on port %v\n", cfg.Env["port"])
	return http.ListenAndServe(fmt.Sprintf(":%v", cfg.Env["port"]), server.routes.Routes())
}
