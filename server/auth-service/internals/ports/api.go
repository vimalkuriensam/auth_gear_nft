package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"

	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
)

type AuthApiPort interface {
	GetUserApi(http.ResponseWriter, *http.Request)
	CreateGRPCUserApi(models.User) pb.AuthResponse
	LoginUserApi(http.ResponseWriter, *http.Request)
	LoginGRPCUserApi(models.User) pb.AuthResponse
	RegisterUserApi(http.ResponseWriter, *http.Request)
	UpdateUserApi(http.ResponseWriter, *http.Request)
	DeleteUserApi(http.ResponseWriter, *http.Request)
}
