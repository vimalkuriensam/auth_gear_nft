package main

import (
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/config"
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/framework/left/queue"
)

func main() {
	configPort := config.Initialize()
	configPort.LoadEnvironment()
	queuePort := queue.Initialize(configPort)
	queuePort.Connect()
	defer configPort.GetConfig().Queue.Connection.Close()
	queuePort.Listen()
}
