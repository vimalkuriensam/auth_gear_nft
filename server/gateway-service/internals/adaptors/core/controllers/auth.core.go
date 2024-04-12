package controllers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

func (authAd Adaptor) CreateUser(input []byte) models.PayloadData {
	payload := models.Payload{
		Id:   uuid.New().String(),
		Kind: "Auth_Register",
		Type: "primary",
		Data: input,
	}
	authAd.config.InitMessages(payload.Id)
	authAd.queue.Emit(payload)
	reply := <-authAd.config.GetConfig().Queue.Messages[payload.Id]
	var payloadData models.PayloadData
	json.Unmarshal(reply.Data, &payloadData)
	return payloadData
}

func (authAd Adaptor) DecodeUser(data []byte) models.UserResponse {
	var resp models.UserResponse
	json.Unmarshal(data, &resp)
	return resp
}
