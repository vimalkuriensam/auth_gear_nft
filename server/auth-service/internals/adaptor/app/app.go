package app

import "net/http"

type Adaptor struct{}

func Initialize() *Adaptor {
	return &Adaptor{}
}

func (appAd *Adaptor) LoginUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) RegisterUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) UpdateUserApi(w http.ResponseWriter, req *http.Request) {}

func (appAd *Adaptor) DeleteUserApi(w http.ResponseWriter, req *http.Request) {}
