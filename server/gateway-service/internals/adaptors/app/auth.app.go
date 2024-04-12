package app

import "net/http"

func (authAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {}

func (authAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {}

func (authAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	authAd.ctrl.CreateUser(input.B)
}

func (authAd *Adaptor) UpdateUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) DeleteUserApi(http.ResponseWriter, *http.Request) {}
