package http

import (
	"github.com/rasouliali1379/terrax/internal/core/domain"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func mapInitDataToUserModel(data initdata.InitData) *domain.User {
	return &domain.User{
		UserID:    uint64(data.User.ID),
		FirstName: data.User.FirstName,
		LastName:  data.User.LastName,
		Username:  data.User.Username,
	}
}
