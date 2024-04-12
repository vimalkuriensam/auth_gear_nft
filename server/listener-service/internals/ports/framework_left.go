package ports

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"

type QueuePort interface {
	Connect()
	Listen()
	Emit(models.Payload) error
}

type RoutesPort interface {
	RouteRequest(models.Payload)
	AuthRoutes(models.Payload, string)
}
