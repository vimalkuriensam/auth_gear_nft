package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"

	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
)

type AuthApiPort interface {
	GetUserApi(http.ResponseWriter, *http.Request)
	GetGRPCUserApi(user models.User) pb.AuthResponse
	CreateGRPCUserApi(models.User) pb.AuthResponse
	LoginGRPCUserApi(models.User) pb.AuthResponse
	UpdateUserApi(http.ResponseWriter, *http.Request)
	DeleteUserApi(http.ResponseWriter, *http.Request)
}
