package ports

import "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"

type AuthApiPort interface {
	CreateUserApi(models.Payload)
}
