package app

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/ports"

type Adaptor struct {
	ctrl ports.ControllersPort
	grpc ports.GRPCPort
}

func Initialize(ctrl ports.ControllersPort, grpc ports.GRPCPort) *Adaptor {
	return &Adaptor{
		ctrl: ctrl,
		grpc: grpc,
	}
}
