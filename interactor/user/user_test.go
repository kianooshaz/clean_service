package user

import (
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/param"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"testing"
)

func TestValidate(t *testing.T) {
	var testCases = []struct {
		user *param.EntryUser
		err  contract.IServiceError
	}{
		{user: &param.EntryUser{
			ID:        0,
			Username:  "kianoosh.az",
			FirstName: "Kianoosh",
			LastName:  "Ashayeri zade",
			Email:     "Kianoosh.az@gmail.com",
			Password:  "123456289515",
			Active:    true,
		}, err: nil},
		{user: &param.EntryUser{
			ID:        0,
			Username:  "kianooshaz",
			FirstName: "kian",
			LastName:  "az",
			Email:     "",
			Password:  "1234589526",
			Active:    false,
		}, err: errors.NewBadRequestError("email is empty")},
		{user: &param.EntryUser{
			ID:        0,
			Username:  "kianooshaz",
			FirstName: "kian",
			LastName:  "az",
			Email:     "kianoosh.com",
			Password:  "1234578965",
			Active:    false,
		}, err: errors.NewBadRequestError("invalid email")},
		{user: &param.EntryUser{
			ID:        0,
			Username:  "kianooshaz",
			FirstName: "kian",
			LastName:  "az",
			Email:     "kianoosh@gmail.com",
			Password:  "",
			Active:    false,
		}, err: errors.NewBadRequestError("password is empty")},
		{user: &param.EntryUser{
			ID:        0,
			Username:  "kianooshaz",
			FirstName: "kian",
			LastName:  "az",
			Email:     "kianoosh@gmail.com",
			Password:  "12345",
			Active:    false,
		}, err: errors.NewBadRequestError("password is less than 8 characters")},
	}
	for _, testCase := range testCases {
		result := validate(testCase.user)
		if result != nil {
			if !result.IsEqual(testCase.err) {
				t.Error(testCase.err, result)
			}
		} else if testCase.err != nil {
			t.Error(testCase.err, result)
		}
	}
}
