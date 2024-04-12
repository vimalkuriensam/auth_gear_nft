package ports

import "github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"

type QueuePort interface {
	Connect()
	Listen()
	Emit(models.Payload) error
}
