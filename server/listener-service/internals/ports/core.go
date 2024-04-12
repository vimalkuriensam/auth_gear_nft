package ports

import "github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/models"

type ConfigPort interface {
	GetConfig() *models.Config
}

type ControllersPort interface {
	ReadUser([]byte) models.User
}
