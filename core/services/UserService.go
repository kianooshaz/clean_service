package services

import (
	"fmt"
	"github.com/kianooshaz/clean_service/core/contract/convertor"
	"github.com/kianooshaz/clean_service/core/contract/interfaces"
	"github.com/kianooshaz/clean_service/core/contract/params"
	"github.com/kianooshaz/clean_service/core/entity"
	"github.com/kianooshaz/clean_service/core/utils/bcrypt"
	"github.com/kianooshaz/clean_service/core/utils/errors"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type userService struct {
	Repo interfaces.IUserRepository
}

func NewUserService(repo interfaces.IUserRepository) interfaces.IUserService {

	return &userService{
		Repo: repo,
	}
}

func (s *userService) Create(entry *params.EntryUser) (*params.PublicUser, interfaces.IServiceError) {

	if serErr := validate(entry); serErr != nil {
		return nil, serErr
	}

	user := convertor.ConvertEntryUserToUser(entry)
	user.Password = bcrypt.GetMd5(user.Password)

	user, serErr := s.Repo.Create(user)
	if serErr != nil {
		return nil, serErr
	}

	return convertor.ConvertUserToPublicUser(user), nil
}

func (s *userService) Get(id int) (*params.PublicUser, interfaces.IServiceError) {

	user, serErr := s.Repo.Get(id)
	if serErr != nil {
		return nil, serErr
	}

	return convertor.ConvertUserToPublicUser(user), nil
}

func (s *userService) Update(entry *params.EntryUser, isPartial bool) (*params.PublicUser, interfaces.IServiceError) {
	fmt.Println("test service") // todo delete
	user, serErr := s.Repo.Get(entry.ID)
	if serErr != nil {
		return nil, serErr
	}

	if isPartial {
		user = partialUpdate(user, entry)
	} else {
		if serErr := validate(entry); serErr != nil {
			return nil, serErr
		}
		user = generalUpdate(user, entry)
	}

	user, serErr = s.Repo.Update(user)
	if serErr != nil {
		return nil, serErr
	}

	return convertor.ConvertUserToPublicUser(user), nil
}

func (s *userService) Delete(id int) interfaces.IServiceError {

	if serErr := s.Repo.Delete(id); serErr != nil {
		return serErr
	}

	return nil
}

func (s *userService) FindAll() ([]params.PublicUser, interfaces.IServiceError) {

	users, serErr := s.Repo.FindAll()
	if serErr != nil {
		return nil, serErr
	}

	var results []params.PublicUser
	for _, user := range users {
		results = append(results, *convertor.ConvertUserToPublicUser(&user))
	}

	return results, nil
}

func validate(user *params.EntryUser) interfaces.IServiceError {
	if user.Email == "" {
		return errors.NewBadRequestError("email is empty")
	}
	if !emailRegex.MatchString(user.Email) {
		return errors.NewBadRequestError("invalid email")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("password is empty")
	}
	if len(user.Password) < 8 {
		return errors.NewBadRequestError("password is less than 8 characters")
	}
	return nil
}

func partialUpdate(user *entity.User, entry *params.EntryUser) *entity.User {
	if entry.Username != "" {
		user.Username = entry.Username
	}
	if entry.FirstName != "" {
		user.FirstName = entry.FirstName
	}
	if entry.LastName != "" {
		user.LastName = entry.LastName
	}
	if entry.Email != "" {
		user.Email = entry.Email
	}
	if entry.Password != "" {
		user.Password = bcrypt.GetMd5(entry.Password)
	}
	return user
}

func generalUpdate(user *entity.User, entry *params.EntryUser) *entity.User {
	user.Username = entry.Username
	user.FirstName = entry.FirstName
	user.LastName = entry.LastName
	user.Email = entry.Email
	user.Password = bcrypt.GetMd5(entry.Password)
	return user
}
