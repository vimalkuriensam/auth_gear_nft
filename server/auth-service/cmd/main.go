package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/app"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/controllers"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/routes"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/left/server"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/framework/right/db"
)

func main() {
	//Initialize Config and Environment
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	//Load the adaptors
	controllerPort := controllers.Initialize(configPort)
	//Initialize DB
	dbPort := db.Initialize(configPort)
	if err := dbPort.DBInit(); err != nil {
		configPort.GetConfig().Logger.Fatalf("unable to initialize db: %v\n", err)
	}
	apiPort := app.Initialize(dbPort, controllerPort)
	routesPort := routes.Initialize(apiPort)
	//Initialize and start the server
	serverPort := server.Initialize(configPort, routesPort)
	serverPort.Server()
}
