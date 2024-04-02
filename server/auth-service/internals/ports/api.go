package ports

import "net/http"

type AuthApiPort interface {
	LoginUserApi(http.ResponseWriter, *http.Request)
	RegisterUserApi(http.ResponseWriter, *http.Request)
	UpdateUserApi(http.ResponseWriter, *http.Request)
	DeleteUserApi(http.ResponseWriter, *http.Request)
}
