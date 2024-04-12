package routes

import (
	"github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/models"
)

func (arAd *Adaptor) AuthRoutes(payload models.Payload, subRequest string) {
	switch subRequest {
	case "Register":
		arAd.api.CreateUserApi(payload)
	}
}
