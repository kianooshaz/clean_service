package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/pkg/claims"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"github.com/kianooshaz/clean_service/pkg/hash"
	"time"
)

type authService struct {
	config config.Config
}

func NewAuthService(config config.Config) contract.IAuthService {
	return &authService{config: config}
}

func (s *authService) GenerateAccessToken(user *entity.User) (string, contract.IServiceError) {
	accessExpirationTime := time.Now().Add(time.Duration(s.config.AccessExpirationInMinute) * time.Minute)

	clm := &claims.Claims{
		Id:             user.ID,
		HashedId:       hash.EncodeId(user.ID),
		Email:          user.Email,
		Username:       user.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: accessExpirationTime.Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, clm)
	tokenString, err := accessToken.SignedString([]byte(s.config.JwtSecret))
	if err != nil {
		return "", errors.NewInternalServerError("jwt error", err)
	}
	return tokenString, nil

}

func (s *authService) GenerateRefreshToken(user *entity.User) (string, contract.IServiceError) {
	refreshExpirationTime := time.Now().Add(time.Duration(s.config.RefreshExpirationInMinute) * time.Minute)

	clm := &claims.Claims{
		Id:             user.ID,
		HashedId:       hash.EncodeId(user.ID),
		Email:          user.Email,
		Username:       user.Email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: refreshExpirationTime.Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, clm)
	tokenString, err := accessToken.SignedString([]byte(s.config.JwtSecret))
	if err != nil {
		return "", errors.NewInternalServerError("jwt error", err)
	}
	return tokenString, nil
}

func (s *authService) GenerateAccountActivationToken(userId uint, email string) (string, contract.IServiceError) {
	panic("implement me")
}

func (s *authService) GeneratePasswordResetToken(userId uint, email string) (string, contract.IServiceError) {
	panic("implement me")
}

func (s *authService) ParseToken(token string) (*claims.Claims, error) {
	panic("implement me")
}
