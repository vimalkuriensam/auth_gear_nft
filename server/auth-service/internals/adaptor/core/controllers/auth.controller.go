package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/models"
	"golang.org/x/crypto/bcrypt"
)

func (cAd *Adaptor) ReadUserRequestController(w http.ResponseWriter, req *http.Request) (models.User, error) {
	value, err := cAd.config.ReadJSON(req)
	if err != nil {
		cAd.config.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return models.User{}, errors.New(err.Error())
	}
	userReq := models.User{}
	_ = json.Unmarshal(value.B, &userReq)
	return userReq, nil
}

func (cAd *Adaptor) LoginController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) GetUserController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) UpdateController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) DeleteController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) PrintRegistration(w http.ResponseWriter, req *http.Request, success bool, status int, data interface{}, msg string) {
	if success {
		cAd.config.WriteJSON(w, http.StatusCreated, data, msg)
	} else {
		cAd.config.ErrorJSON(w, req.URL.Path, msg, status)
	}
}

func (cAd *Adaptor) PaswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (cAd *Adaptor) ComparePassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
