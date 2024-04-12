package app

import (
	"fmt"

	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/framework/left/queue"
)

func (appAd *Adaptor) CreateUserApi(payload models.Payload) {
	user := appAd.ctrl.ReadUser(payload.Data)
	respBytes, _ := appAd.grpc.RegisterUser(user)
	queueAdaptor := queue.GetAdaptor()
	err := queueAdaptor.Emit(models.Payload{
		Id:   payload.Id,
		Kind: "Auth_Register",
		Type: payload.Type,
		Data: respBytes,
	})
	fmt.Println("error 2:", err)
}
