package app

import (
	"net/http"
)

func (authAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {}

func (authAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {}

func (authAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	resp := authAd.ctrl.CreateUser(input.B)
	if !resp.Success {
		authAd.config.ErrorJSON(w, req.URL.Path, resp.Message, resp.Code)
		return
	}
	data := authAd.ctrl.DecodeUser(resp.Data)
	authAd.config.WriteJSON(w, resp.Code, data, resp.Message)
}

func (authAd *Adaptor) UpdateUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) DeleteUserApi(http.ResponseWriter, *http.Request) {}
