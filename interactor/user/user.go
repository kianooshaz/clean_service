package user

import (
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/param"
	"github.com/kianooshaz/clean_service/pkg/bcrypt"
	"github.com/kianooshaz/clean_service/pkg/errors"
)

type userService struct {
	repo     contract.IUserRepository
	config   config.Config
	auth     contract.IAuthService
	validate contract.IUserValidation
}

func NewService(c config.Config, r contract.IUserRepository, a contract.IAuthService, v contract.IUserValidation) contract.IUserService {

	return &userService{
		repo:     r,
		config:   c,
		auth:     a,
		validate: v,
	}
}

func (s *userService) Create(entry *param.EntryUser) (*param.PublicUser, contract.IServiceError) {
	if serErr := s.validate.EmailAndPasswordValidation(entry); serErr != nil {
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
		if serErr := s.validate.EmailAndPasswordValidation(entry); serErr != nil {
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

func (s *userService) Login(r *param.LoginRequestUser) (*param.UserTokens, contract.IServiceError) {
	const section = errors.Section("Interactor.user.Login")

	if serErr := s.validate.LoginValidation(r); serErr != nil {
		return nil, serErr
	}

	entryUser := &param.EntryUser{
		Email:    r.Email,
		Password: r.Password,
	}
	user, serErr := s.repo.GetUserByEmail(entryUser.Email)
	if serErr != nil {
		return nil, serErr
	}

	if user.Password != bcrypt.GetMd5(entryUser.Password, s.config.Salt) {
		return nil, errors.NewUnauthorizedError(section, "incorrect password")
	}

	accessToken, serErr := s.auth.GenerateAccessToken(user)
	if serErr != nil {
		return nil, serErr
	}

	refreshToken, serErr := s.auth.GenerateRefreshToken(user)
	if serErr != nil {
		return nil, serErr
	}

	tokens := &param.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
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
