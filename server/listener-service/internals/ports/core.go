package ports

import (
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"google.golang.org/grpc"
)

type ConfigPort interface {
	GetConfig() *models.Config
}

type ControllersPort interface {
	ReadUser([]byte) models.User
}

type GRPCPort interface {
	DialAuth() (*grpc.ClientConn, error)
	RegisterUser(user models.User) ([]byte, error)
}
