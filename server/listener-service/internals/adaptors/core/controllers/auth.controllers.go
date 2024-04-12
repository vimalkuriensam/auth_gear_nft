package controllers

import (
	"encoding/json"

	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
)

func (authAd *Adaptor) ReadUser(data []byte) models.User {
	var user models.User
	_ = json.Unmarshal(data, &user)
	return user
}
