package controllers

import (
	"encoding/json"

	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/ports"
)

type Adaptor struct {
	config ports.ConfigPort
	queue  ports.QueuePort
}

func Initialize(config ports.ConfigPort, queue ports.QueuePort) *Adaptor {
	return &Adaptor{
		config: config,
		queue:  queue,
	}
}

func (bcCtrl *Adaptor) EmitData(payload models.Payload) models.PayloadData {
	bcCtrl.config.InitMessages(payload.Id)
	bcCtrl.queue.Emit(payload)
	reply := <-bcCtrl.config.GetConfig().Queue.Messages[payload.Id]
	var payloadData models.PayloadData
	json.Unmarshal(reply.Data, &payloadData)
	return payloadData
}
