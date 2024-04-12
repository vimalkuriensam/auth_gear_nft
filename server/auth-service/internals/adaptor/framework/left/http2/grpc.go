package http2

import (
	"context"
	"log"
	"net"

	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"
	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/ports"
	"google.golang.org/grpc"
)

type Adaptor struct {
	app ports.AuthApiPort
	pb.AuthenticationsServer
}

func Initialize(app ports.AuthApiPort) *Adaptor {
	return &Adaptor{
		app: app,
	}
}

func (grpcAd *Adaptor) Listen() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s\n", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterAuthenticationsServer(s, grpcAd)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func (grpcAd *Adaptor) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	resp := grpcAd.app.CreateGRPCUserApi(user)
	return &resp, nil
}

func (grpcAd *Adaptor) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	resp := grpcAd.app.LoginGRPCUserApi(user)
	return &resp, nil
}
