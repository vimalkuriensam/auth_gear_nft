package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/app"
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/controllers"
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/framework/left/queue"
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/framework/left/routes"
)

func main() {
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	controllersPort := controllers.Initialize()
	apiPort := app.Initialize(controllersPort)
	routesPort := routes.Initialize(apiPort)
	queuePort := queue.Initialize(configPort, routesPort)
	queuePort.Connect()
	defer configPort.GetConfig().Queue.Connection.Close()
	queuePort.Listen()
}
