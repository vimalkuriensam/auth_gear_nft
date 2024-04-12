package controllers

import (
	"github.com/google/uuid"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

func (authAd Adaptor) CreateUser(input []byte) {
	payload := models.Payload{
		Id:   uuid.New().String(),
		Kind: "Auth_Register",
		Type: "primary",
		Data: input,
	}
	authAd.config.InitMessages(payload.Id)
	authAd.queue.Emit(payload)
	// reply := <- authAd.config.GetConfig().Queue.Messages[payload.Id]
}
