package app

import "net/http"

func (authAd *Adaptor) GetUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) LoginUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) RegisterUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) UpdateUserApi(http.ResponseWriter, *http.Request) {}

func (authAd *Adaptor) DeleteUserApi(http.ResponseWriter, *http.Request) {}
