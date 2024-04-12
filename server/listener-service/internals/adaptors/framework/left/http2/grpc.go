package http2

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/ports"

type Adaptor struct {
	config ports.ConfigPort
}

func Initialize(config ports.ConfigPort) *Adaptor {
	return &Adaptor{
		config: config,
	}
}
