package app

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/ports"

type Adaptor struct {
	ctrl ports.ControllersPort
}

func Initialize(ctrl ports.ControllersPort) *Adaptor {
	return &Adaptor{
		ctrl: ctrl,
	}
}
