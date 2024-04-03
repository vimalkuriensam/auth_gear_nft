package controllers

import "github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"

type Adaptor struct {
	config ports.ConfigPort
}

func Initialize(cfg ports.ConfigPort) *Adaptor {
	return &Adaptor{
		config: cfg,
	}
}
