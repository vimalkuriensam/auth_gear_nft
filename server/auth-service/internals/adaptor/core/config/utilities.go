package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/models"
)

func (cfgAd *Adaptor) ReadJSON(req *http.Request) (models.ReadValue, error) {
	data := models.ReadValue{
		B: []byte(""),
		D: nil,
	}
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err = json.Unmarshal([]byte(body), &data.D); err == nil {
			data.B = []byte(body)
		} else {
			return models.ReadValue{}, err
		}
	} else {
		return models.ReadValue{}, err
	}
	return data, nil
}

func (cfgAd *Adaptor) WriteJSON(w http.ResponseWriter, status int, data interface{}, msg string, headers ...http.Header) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	cfgAd.config.Response.Data = data
	cfgAd.config.Response.Message = msg
	if b_data, err := json.Marshal(cfg.Response); err == nil {
		w.Write(b_data)
	}
}

func (cfgAd *Adaptor) ErrorJSON(w http.ResponseWriter, path string, reason string, status ...int) {
	errorStatus := http.StatusBadRequest
	if len(status) > 0 {
		errorStatus = status[0]
	}
	cfgAd.config.Logger.Println("error-reason: ", reason)
	cfgAd.config.Error = &models.ErrorResponse{
		Status:    errorStatus,
		Path:      path,
		Message:   reason,
		Timestamp: time.Now(),
	}
	cfgAd.WriteJSON(w, errorStatus, cfg.Error, "Error")
}
