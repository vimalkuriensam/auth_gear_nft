package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/ports"
	"github.com/vimalkuriensam/auto_gear_nft/auth-service/pkg/constants"

	pb "github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/framework/left/http2/proto"
)

type Adaptor struct {
	config     ports.ConfigPort
	db         ports.DBPort
	controller ports.AuthController
}

func Initialize(config ports.ConfigPort, db ports.DBPort, ctrl ports.AuthController) *Adaptor {
	return &Adaptor{
		config:     config,
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

func (appAd *Adaptor) GetGRPCUserApi(user models.User) pb.AuthResponse {
	user, err := appAd.db.GetUserByID(user.ID)
	if err != nil {
		return appAd.config.ErrorResponse(constants.USER_NONEXIST, http.StatusBadRequest)
	}
	user.Password = ""
	bt, _ := json.Marshal(user)
	return appAd.config.SuccessResponse(constants.USER_FETCH_SUCCESS, http.StatusOK, bt)
}

func (appAd *Adaptor) CreateGRPCUserApi(user models.User) pb.AuthResponse {
	hash, err := appAd.controller.PaswordHash(user.Password)
	if err != nil {
		return appAd.config.ErrorResponse(constants.PASSWORD_HASH_ERROR, http.StatusInternalServerError)
	}
	user.Password = string(hash)
	if inserted_data, err := appAd.db.InsertUser(user); err == nil {
		token, err := appAd.controller.GenerateJWTToken(inserted_data)
		if err != nil {
			return appAd.config.ErrorResponse(constants.TOKEN_GENERATION_ERROR, http.StatusInternalServerError)
		}
		var responseData models.UserResponse
		inserted_data.Password = ""
		responseData.User = inserted_data
		responseData.Token = token
		bt, _ := json.Marshal(responseData)
		return appAd.config.SuccessResponse(constants.REGISTRATION_SUCCESS, http.StatusCreated, bt)
	} else {
		return appAd.config.ErrorResponse(constants.DATA_INSERTION_ERROR, http.StatusInternalServerError)
	}
}

func (appAd *Adaptor) LoginGRPCUserApi(user models.User) pb.AuthResponse {
	existingUser, err := appAd.db.GetUserByEmail(user.Email)
	if err != nil {
		return appAd.config.ErrorResponse(constants.INVALID_USER_ERROR, http.StatusBadRequest)
	}
	isPasswordMatch := appAd.controller.ComparePassword(existingUser.Password, user.Password)
	if !isPasswordMatch {
		return appAd.config.ErrorResponse(constants.INVALID_USER_ERROR, http.StatusBadRequest)
	}
	token, err := appAd.controller.GenerateJWTToken(existingUser)
	if err != nil {
		return appAd.config.ErrorResponse(constants.TOKEN_GENERATION_ERROR, http.StatusInternalServerError)
	}
	existingUser.Password = ""
	var userResponse models.UserResponse
	userResponse.User = existingUser
	userResponse.Token = token
	bt, _ := json.Marshal(userResponse)
	return appAd.config.SuccessResponse(constants.LOGIN_SUCCESS, http.StatusCreated, bt)
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
