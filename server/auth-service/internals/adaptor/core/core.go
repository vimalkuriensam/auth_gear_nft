package core

import "net/http"

type Adaptor struct{}

func Initialize() *Adaptor {
	return &Adaptor{}
}

func (cAd *Adaptor) RegisterController(w http.ResponseWriter, req *http.Request) {

}

func (cAd *Adaptor) LoginController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) GetUserController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) UpdateController(w http.ResponseWriter, req *http.Request) {}

func (cAd *Adaptor) DeleteController(w http.ResponseWriter, req *http.Request) {}
