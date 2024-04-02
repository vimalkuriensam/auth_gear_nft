package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auth_gear_nft/auth-service/internals/adaptor/core/models"
)

type AuthController interface {
	RegisterController(http.ResponseWriter, *http.Request)
	LoginController(http.ResponseWriter, *http.Request)
	GetUserController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
}

type ConfigPort interface {
	GetConfig() *models.Config
	LoadEnvironment() error
	ReadJSON(*http.Request) (models.ReadValue, error)
	WriteJSON(http.ResponseWriter, int, interface{}, string, ...http.Header)
}
