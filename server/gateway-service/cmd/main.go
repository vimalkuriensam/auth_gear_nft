package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/app"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/controllers"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/framework/left/routes"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/framework/left/server"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/framework/right/queue"
)

func main() {
	// Initialize config adaptor
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	//initialize, connect and listen to queue messages
	queuePort := queue.Initialize(configPort)
	queuePort.Connect()
	go queuePort.Listen()
	controllersPort := controllers.Initialize(configPort, queuePort)
	// Initialize api adaptor
	apiPort := app.Initialize(configPort, controllersPort)
	routesPort := routes.Initialize(apiPort)
	serverPort := server.Initialize(configPort, routesPort)
	// Start the server
	configPort.GetConfig().Logger.Fatal(serverPort.Serve())
}
