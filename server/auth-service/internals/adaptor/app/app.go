package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/ports"
)

type Adaptor struct {
	db         ports.DBPort
	controller ports.AuthController
}

func Initialize(db ports.DBPort, ctrl ports.AuthController) *Adaptor {
	return &Adaptor{
		controller: ctrl,
		db:         db,
	}
}

func (appAd *Adaptor) GetUserApi(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	id_num, err := strconv.Atoi(id)
	if err != nil {
		appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, err.Error())
		return
	}
	user, err := appAd.db.GetUserByID(uint(id_num))
	if err != nil {
		appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
		return
	}
	appAd.controller.PrintRegistration(w, req, true, http.StatusOK, user, "User Fetched")
}

func (appAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	user_data, err := appAd.controller.ReadUserRequestController(w, req)
	if err == nil {
		if inserted_data, err := appAd.db.InsertUser(user_data); err == nil {
			appAd.controller.PrintRegistration(w, req, true, http.StatusCreated, inserted_data, "User Created")
		} else {
			fmt.Println(err.Error())
			appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
		}
	} else {
		appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
	}
}

func (appAd *Adaptor) UpdateUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) DeleteUserApi(w http.ResponseWriter, req *http.Request) {}
