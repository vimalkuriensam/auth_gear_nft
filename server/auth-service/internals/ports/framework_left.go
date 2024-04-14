package ports

import (
	"context"
	"net/http"

	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
)

type RoutesPort interface {
	Routes() http.Handler
}

type GRPCPort interface {
	Listen()
	Register(context.Context, *pb.RegisterRequest) (*pb.AuthResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.AuthResponse, error)
	GetUser(context.Context, *pb.GetIDRequest) (*pb.AuthResponse, error)
	DeleteUser(context.Context, *pb.GetIDRequest) (*pb.AuthResponse, error)
}
