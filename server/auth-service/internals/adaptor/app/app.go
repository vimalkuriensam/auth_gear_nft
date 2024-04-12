package app

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/ports"
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
	user.Password = ""
	appAd.controller.PrintRegistration(w, req, true, http.StatusOK, user, "User Fetched")
}

func (appAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {
	user_data, err := appAd.controller.ReadUserRequestController(w, req)
	if err == nil {
		user, err := appAd.db.GetUserByEmail(user_data.Email)
		if err != nil {
			appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, "invalid credentials")
			return
		}
		isPasswordMatch := appAd.controller.ComparePassword(user.Password, user_data.Password)
		if !isPasswordMatch {
			appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, "invalid credentials")
			return
		}
		token, err := appAd.controller.GenerateJWTToken(user)
		if err != nil {
			appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
		}
		user.Password = ""
		var response models.UserResponse
		response.User = user
		response.Token = token
		appAd.controller.PrintRegistration(w, req, true, http.StatusOK, response, "User login successful")
	} else {
		appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
	}
}

func (appAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {
	user_data, err := appAd.controller.ReadUserRequestController(w, req)
	if err == nil {
		hash, err := appAd.controller.PaswordHash(user_data.Password)
		if err != nil {
			appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
			return
		}
		user_data.Password = string(hash)
		if inserted_data, err := appAd.db.InsertUser(user_data); err == nil {
			token, err := appAd.controller.GenerateJWTToken(inserted_data)
			if err != nil {
				appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
			}
			var responseData models.UserResponse
			inserted_data.Password = ""
			responseData.User = inserted_data
			responseData.Token = token
			appAd.controller.PrintRegistration(w, req, true, http.StatusCreated, responseData, "User Created")
		} else {
			appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
		}
	} else {
		appAd.controller.PrintRegistration(w, req, false, http.StatusInternalServerError, nil, err.Error())
	}
}

func (appAd *Adaptor) UpdateUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) DeleteUserApi(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	id_num, err := strconv.Atoi(id)
	if err != nil {
		appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, err.Error())
		return
	}
	user, err := appAd.db.GetUserByID(uint(id_num))
	if err != nil {
		appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, "user does not exist")
		return
	}
	if err = appAd.db.DeleteUserByID(uint(id_num)); err != nil {
		appAd.controller.PrintRegistration(w, req, false, http.StatusBadRequest, nil, "error deleting user")
		return
	}
	user.Password = ""
	appAd.controller.PrintRegistration(w, req, true, http.StatusCreated, user, "User Deleted")
}
