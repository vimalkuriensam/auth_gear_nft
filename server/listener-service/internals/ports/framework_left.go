package ports

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"

type QueuePort interface {
	Connect()
	Listen()
}

type RoutesPort interface {
	RouteRequest(models.Payload)
	AuthRoutes(models.Payload, string)
}
