package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"

	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
)

type AuthApiPort interface {
	GetUserApi(http.ResponseWriter, *http.Request)
	CreateGRPCUserApi(models.User) pb.RegisterResponse
	LoginUserApi(http.ResponseWriter, *http.Request)
	RegisterUserApi(http.ResponseWriter, *http.Request)
	UpdateUserApi(http.ResponseWriter, *http.Request)
	DeleteUserApi(http.ResponseWriter, *http.Request)
}
