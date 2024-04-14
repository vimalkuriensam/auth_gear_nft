package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/pkg/enums"
)

func (authAd *Adaptor) GetUserId(req *http.Request) ([]byte, error) {
	id := chi.URLParam(req, "id")
	var user models.User
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return []byte{}, err
	}
	user.ID = uint(idInt)
	input, _ := json.Marshal(user)
	return input, nil
}

func (authAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {
	input, err := authAd.GetUserId(req)
	if err != nil {
		authAd.config.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	resp := authAd.ctrl.GetUser(input)
	authAd.SendResponse(w, req, resp, enums.UserWithoutToken)
}

func (authAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	resp := authAd.ctrl.LoginUser(input.B)
	authAd.SendResponse(w, req, resp, enums.UserWithToken)
}

func (authAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	input, _ := authAd.config.ReadJSON(req)
	resp := authAd.ctrl.CreateUser(input.B)
	authAd.SendResponse(w, req, resp, enums.UserWithToken)
}

func (authAd *Adaptor) UpdateUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) DeleteUserApi(w http.ResponseWriter, req *http.Request) {
	input, err := authAd.GetUserId(req)
	if err != nil {
		authAd.config.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	resp := authAd.ctrl.DeleteUser(input)
	authAd.SendResponse(w, req, resp, enums.UserWithoutToken)
}

func (authAd *Adaptor) SendResponse(w http.ResponseWriter, req *http.Request, resp models.PayloadData, kind enums.UserKind) {
	if !resp.Success {
		authAd.config.ErrorJSON(w, req.URL.Path, resp.Message, resp.Code)
		return
	}
	var data any
	if kind == enums.UserWithToken {
		data = authAd.ctrl.DecodeUserResponse(resp.Data)
	} else {
		data = authAd.ctrl.DecodeUser(resp.Data)
	}
	authAd.config.WriteJSON(w, resp.Code, data, resp.Message)
}
