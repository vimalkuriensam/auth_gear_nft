package ports

import (
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"google.golang.org/grpc"
)

type QueuePort interface {
	Connect()
	Listen()
	Emit(models.Payload) error
}

type RoutesPort interface {
	RouteRequest(models.Payload)
	AuthRoutes(models.Payload, string)
}

type GRPCPort interface {
	DialAuth() (*grpc.ClientConn, error)
	RegisterUser(models.User) ([]byte, error)
	LoginUser(models.User) ([]byte, error)
	GetUser(models.User) ([]byte, error)
	DeleteUser(models.User) ([]byte, error)
}
