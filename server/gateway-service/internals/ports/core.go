package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

type ConfigPort interface {
	GetConfig() *models.Config
	LoadEnvironment() error
	ReadJSON(*http.Request) (models.ReadValue, error)
	WriteJSON(http.ResponseWriter, int, interface{}, string, ...http.Header)
	ErrorJSON(http.ResponseWriter, string, string, ...int)
	InitMessages(string)
	SetMessage(string, models.Payload)
	DeleteMessage(string)
}

type ControllersPort interface {
	GetPayload() models.Payload
	CreateUser([]byte) models.PayloadData
	LoginUser([]byte) models.PayloadData
	GetUser([]byte) models.PayloadData
	DeleteUser([]byte) models.PayloadData
	DecodeUserResponse([]byte) models.UserResponse
	DecodeUser([]byte) models.User
	EmitData(models.Payload) models.PayloadData
}
