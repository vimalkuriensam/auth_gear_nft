package app

import (
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/framework/left/queue"
)

func (appAd *Adaptor) CreateUserApi(payload models.Payload) {
	user := appAd.ctrl.ReadUser(payload.Data)
	respBytes, _ := appAd.grpc.RegisterUser(user)
	queueAdaptor := queue.GetAdaptor()
	_ = queueAdaptor.Emit(models.Payload{
		Id:   payload.Id,
		Kind: "Auth_Register",
		Type: payload.Type,
		Data: respBytes,
	})
}

func (appAd *Adaptor) LoginUserApi(payload models.Payload) {
	user := appAd.ctrl.ReadUser(payload.Data)
	respBytes, _ := appAd.grpc.LoginUser(user)
	queueAdaptor := queue.GetAdaptor()
	queueAdaptor.Emit(models.Payload{
		Id:   payload.Id,
		Kind: "Auth_Login",
		Type: payload.Type,
		Data: respBytes,
	})
}

func (appAd *Adaptor) GetUserApi(payload models.Payload) {
	user := appAd.ctrl.ReadUser(payload.Data)
	respBytes, _ := appAd.grpc.GetUser(user)
	queueAdaptor := queue.GetAdaptor()
	queueAdaptor.Emit(models.Payload{
		Id:   payload.Id,
		Kind: "Auth_GetUser",
		Type: payload.Type,
		Data: respBytes,
	})
}
