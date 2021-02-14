package user

import (
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/param"
	"github.com/kianooshaz/clean_service/pkg/bcrypt"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type userService struct {
	repo   contract.IUserRepository
	config config.Config
}

func (s *userService) Create(entry *param.EntryUser) (*param.PublicUser, contract.IServiceError) {

	if serErr := validate(entry); serErr != nil {
		return nil, serErr
	}

	user := param.ConvertEntryUserToUser(entry)
	user.Password = bcrypt.GetMd5(user.Password, s.config.Salt)

	user, serErr := s.repo.CreateUser(user)
	if serErr != nil {
		return nil, serErr
	}

	return param.ConvertUserToPublicUser(user), nil
}

func (s *userService) Get(id int) (*param.PublicUser, contract.IServiceError) {

	user, serErr := s.repo.GetUserByID(id)
	if serErr != nil {
		return nil, serErr
	}

	return param.ConvertUserToPublicUser(user), nil
}

func (s *userService) Update(entry *param.EntryUser, isPartial bool) (*param.PublicUser, contract.IServiceError) {
	user, serErr := s.repo.GetUserByID(entry.ID)
	if serErr != nil {
		return nil, serErr
	}

	if isPartial {
		user = partialUpdate(user, entry)
		if entry.Password != "" {
			user.Password = bcrypt.GetMd5(entry.Password, s.config.Salt)
		}
	} else {
		if serErr := validate(entry); serErr != nil {
			return nil, serErr
		}
		user = generalUpdate(user, entry)
		user.Password = bcrypt.GetMd5(entry.Password, s.config.Salt)
	}

	user, serErr = s.repo.UpdateUser(user)
	if serErr != nil {
		return nil, serErr
	}

	return param.ConvertUserToPublicUser(user), nil
}

func (s *userService) Delete(id int) contract.IServiceError {

	if serErr := s.repo.DeleteUserByID(id); serErr != nil {
		return serErr
	}

	return nil
}

func (s *userService) FindAll() ([]param.PublicUser, contract.IServiceError) {

	users, serErr := s.repo.FindAllUser()
	if serErr != nil {
		return nil, serErr
	}

	var results []param.PublicUser
	for _, user := range users {
		results = append(results, *param.ConvertUserToPublicUser(&user))
	}

	return results, nil
}

func validate(user *param.EntryUser) contract.IServiceError {
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

func partialUpdate(user *entity.User, entry *param.EntryUser) *entity.User {
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
	return user
}

func generalUpdate(user *entity.User, entry *param.EntryUser) *entity.User {
	user.Username = entry.Username
	user.FirstName = entry.FirstName
	user.LastName = entry.LastName
	user.Email = entry.Email
	return user
}
