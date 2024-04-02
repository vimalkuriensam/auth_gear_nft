package ports

import "net/http"

type AuthController interface {
	RegisterController(http.ResponseWriter, *http.Request)
	LoginController(http.ResponseWriter, *http.Request)
	GetUserController(http.ResponseWriter, *http.Request)
	UpdateController(http.ResponseWriter, *http.Request)
	DeleteController(http.ResponseWriter, *http.Request)
}
