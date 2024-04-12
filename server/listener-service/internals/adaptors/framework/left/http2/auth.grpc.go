package http2

import (
	"errors"
	"fmt"

	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
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

func (grpcAd *Adaptor) RegisterUser(user models.User) {}
