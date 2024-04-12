package ports

import "github.com/vimalkuriensam/auth_gear_nft/listener-service/internals/adaptors/core/models"

type AuthApiPort interface {
	CreateUserApi(models.Payload)
}
