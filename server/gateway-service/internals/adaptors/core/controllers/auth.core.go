package controllers

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

func (authAd Adaptor) GetPayload() models.Payload {
	return models.Payload{
		Id:   uuid.New().String(),
		Kind: "",
		Type: "primary",
		Data: []byte{},
	}
}

func (authAd Adaptor) CreateUser(input []byte) models.PayloadData {
	payload := authAd.GetPayload()
	payload.Kind = "Auth_Register"
	payload.Data = input
	return authAd.EmitData(payload)
}

func (authAd Adaptor) LoginUser(input []byte) models.PayloadData {
	payload := authAd.GetPayload()
	payload.Kind = "Auth_Login"
	payload.Data = input
	return authAd.EmitData(payload)
}

func (authAd *Adaptor) GetUser(input []byte) models.PayloadData {
	payload := authAd.GetPayload()
	payload.Kind = "Auth_GetUser"
	payload.Data = input
	return authAd.EmitData(payload)
}

func (authAd Adaptor) DecodeUserResponse(data []byte) models.UserResponse {
	var resp models.UserResponse
	json.Unmarshal(data, &resp)
	return resp
}

func (authAd Adaptor) DecodeUser(data []byte) models.User {
	var resp models.User
	json.Unmarshal(data, &resp)
	return resp
}
