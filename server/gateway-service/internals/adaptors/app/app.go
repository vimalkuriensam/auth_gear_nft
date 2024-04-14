package app

import "github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/ports"

type Adaptor struct {
	config ports.ConfigPort
	ctrl   ports.ControllersPort
}

func Initialize(config ports.ConfigPort, ctrl ports.ControllersPort) *Adaptor {
	return &Adaptor{
		config: config,
		ctrl:   ctrl,
	}
}
