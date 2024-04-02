package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/routes"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/server"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/pkg/config"
)

func main() {
	// Initialize the config directory
	cfg := config.Initialize()
	cfg.LoadEnvironment()
	//Load the adaptors
	routesPort := routes.Initialize()
	serverPort := server.Initialize(routesPort)
	serverPort.Server()
}
