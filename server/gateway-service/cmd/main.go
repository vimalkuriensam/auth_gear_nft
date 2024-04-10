package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/app"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/framework/left/routes"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/framework/left/server"
)

func main() {
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	apiPort := app.Initialize()
	routesPort := routes.Initialize(apiPort)
	serverPort := server.Initialize(configPort, routesPort)
	configPort.GetConfig().Logger.Fatal(serverPort.Serve())
}
