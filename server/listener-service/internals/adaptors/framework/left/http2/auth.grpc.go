package http2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	pb "github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/framework/left/http2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (grpcAd *Adaptor) DialAuth() (*grpc.ClientConn, error) {
	config := grpcAd.config.GetConfig()
	env := config.Env
	host, ok1 := env["auth_grpc_host"].(string)
	port, ok2 := env["auth_grpc_port"].(string)
	if !ok1 || !ok2 {
		return nil, errors.New("authentication grpc host or port not set")
	}
	authURL := fmt.Sprintf("%s:%s", host, port)
	return grpc.Dial(authURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func (grpcAd *Adaptor) RegisterUser(user models.User) ([]byte, error) {
	conn, err := grpcAd.DialAuth()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to connect to auth service: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthenticationsClient(conn)
	input := &pb.RegisterRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Email:     user.Email,
	}
	res, err := c.Register(context.Background(), input)
	if err != nil {
		return []byte{}, fmt.Errorf("error registering user: %s", err.Error())
	}
	return json.Marshal(res)
}

func (grpcAd *Adaptor) LoginUser(user models.User) ([]byte, error) {
	conn, err := grpcAd.DialAuth()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to connect to auth service: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthenticationsClient(conn)
	input := &pb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := c.Login(context.Background(), input)
	if err != nil {
		return []byte{}, fmt.Errorf("error logging user: %s", err.Error())
	}
	return json.Marshal(res)
}

func (grpcAd *Adaptor) GetUser(user models.User) ([]byte, error) {
	conn, err := grpcAd.DialAuth()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to connect to auth service: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewAuthenticationsClient(conn)
	input := &pb.GetUserRequest{
		Id: user.Id,
	}
	res, err := c.GetUser(context.Background(), input)
	if err != nil {
		return []byte{}, fmt.Errorf("error getting user: %s", err.Error())
	}
	return json.Marshal(res)
}
