package convertor

import (
	"github.com/kianooshaz/clean_service/core/contract/param"
	"github.com/kianooshaz/clean_service/core/entity"
)

func ConvertUserToPublicUser(user *entity.User) *param.PublicUser {
	return &param.PublicUser{
		ID:        int(user.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Active:    user.Active,
	}
}

func ConvertEntryUserToUser(entry *param.EntryUser) *entity.User {
	return &entity.User{
		Base:      entity.Base{ID: uint(entry.ID)},
		Username:  entry.Username,
		FirstName: entry.FirstName,
		LastName:  entry.LastName,
		Email:     entry.Email,
		Password:  entry.Password,
		Active:    entry.Active,
	}
}
