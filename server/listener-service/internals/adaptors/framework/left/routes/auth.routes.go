package routes

import (
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
)

func (arAd *Adaptor) AuthRoutes(payload models.Payload, subRequest string) {
	switch subRequest {
	case "Register":
		arAd.api.CreateUserApi(payload)
	case "Login":
		arAd.api.LoginUserApi(payload)
	case "GetUser":
		arAd.api.GetUserApi(payload)
	case "DeleteUser":

	}
}
