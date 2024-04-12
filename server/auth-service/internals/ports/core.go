package ports

import (
	"net/http"

	"github.com/vimalkuriensam/auto_gear_nft/auth-service/internals/adaptor/core/models"
)

type AuthController interface {
	ReadUserRequestController(http.ResponseWriter, *http.Request) (models.User, error)
	LoginController(http.ResponseWriter, *http.Request)
	GetUserController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
	PrintRegistration(http.ResponseWriter, *http.Request, bool, int, interface{}, string)
	PaswordHash(password string) ([]byte, error)
	ComparePassword(hash, password string) bool
	GenerateJWTToken(models.User) (string, error)
}

type ConfigPort interface {
	GetConfig() *models.Config
	LoadEnvironment() error
	ReadJSON(*http.Request) (models.ReadValue, error)
	WriteJSON(http.ResponseWriter, int, interface{}, string, ...http.Header)
	ErrorJSON(http.ResponseWriter, string, string, ...int)
}
