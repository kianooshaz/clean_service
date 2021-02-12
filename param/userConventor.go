package param

import (
	"github.com/kianooshaz/clean_service/entity"
)

func ConvertUserToPublicUser(user *entity.User) *PublicUser {
	return &PublicUser{
		ID:        int(user.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Active:    user.Active,
	}
}

func ConvertEntryUserToUser(entry *EntryUser) *entity.User {
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
