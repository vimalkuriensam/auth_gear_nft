package app

import (
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"
)

type Adaptor struct {
	controller ports.AuthController
}

func Initialize(ctrl ports.AuthController) *Adaptor {
	return &Adaptor{
		controller: ctrl,
	}
}

func (appAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {

}

func (appAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	appAd.controller.RegisterController(w, req)
}

func (appAd *Adaptor) UpdateUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) DeleteUserApi(w http.ResponseWriter, req *http.Request) {}
