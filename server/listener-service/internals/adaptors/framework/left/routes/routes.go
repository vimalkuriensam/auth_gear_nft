package routes

import (
	"strings"

	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/ports"
)

type Adaptor struct {
	api ports.AuthApiPort
}

func Initialize(api ports.AuthApiPort) *Adaptor {
	return &Adaptor{
		api: api,
	}
}

func (arAd *Adaptor) RouteRequest(payload models.Payload) {
	splitStrings := strings.Split(payload.Kind, "_")
	switch splitStrings[0] {
	case "Auth":
		arAd.AuthRoutes(payload, splitStrings[1])
	}
}
