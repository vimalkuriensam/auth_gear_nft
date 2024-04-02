package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/app"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/controllers"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/routes"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/server"
)

func main() {
	//Load the adaptors
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	controllerPort := controllers.Initialize()
	apiPort := app.Initialize(controllerPort)
	routesPort := routes.Initialize(apiPort)
	serverPort := server.Initialize(configPort, routesPort)
	serverPort.Server()
}
