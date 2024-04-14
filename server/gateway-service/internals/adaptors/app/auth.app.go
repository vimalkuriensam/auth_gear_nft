package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
)

func (authAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	var user models.User
	idInt, err := strconv.Atoi(id)
	if err != nil {
		authAd.config.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = uint(idInt)
	input, _ := json.Marshal(user)
	resp := authAd.ctrl.GetUser(input)
	if !resp.Success {
		authAd.config.ErrorJSON(w, req.URL.Path, resp.Message, resp.Code)
		return
	}
	data := authAd.ctrl.DecodeUser(resp.Data)
	authAd.config.WriteJSON(w, resp.Code, data, resp.Message)
}

func (authAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	resp := authAd.ctrl.LoginUser(input.B)
	if !resp.Success {
		authAd.config.ErrorJSON(w, req.URL.Path, resp.Message, resp.Code)
		return
	}
	data := authAd.ctrl.DecodeUserResponse(resp.Data)
	authAd.config.WriteJSON(w, resp.Code, data.User, resp.Message)
}

func (authAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	resp := authAd.ctrl.CreateUser(input.B)
	if !resp.Success {
		authAd.config.ErrorJSON(w, req.URL.Path, resp.Message, resp.Code)
		return
	}
	data := authAd.ctrl.DecodeUserResponse(resp.Data)
	authAd.config.WriteJSON(w, resp.Code, data, resp.Message)
}

func (authAd *Adaptor) UpdateUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) DeleteUserApi(http.ResponseWriter, *http.Request) {}
