package server

import (
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/pkg/config"
)

type Adaptor struct {
	routes ports.RoutesPort
}

func Initialize(routes ports.RoutesPort) *Adaptor {
	return &Adaptor{
		routes: routes,
	}
}

func (server *Adaptor) Server() error {
	cfg := config.GetConfig()
	cfg.Logger.Printf("Server is running on port %v\n", cfg.Env["port"])
	return http.ListenAndServe(fmt.Sprintf(":%v", cfg.Env["port"]), server.routes.Routes())
}
