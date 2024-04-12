package controllers

import "github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/ports"

type Adaptor struct {
	config ports.ConfigPort
	queue  ports.QueuePort
}

func Initialize(config ports.ConfigPort, queue ports.QueuePort) *Adaptor {
	return &Adaptor{
		config: config,
	}
}
