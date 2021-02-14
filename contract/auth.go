package contract

import (
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/pkg/claims"
)

type IAuthService interface {
	GenerateAccessToken(user *entity.User) (string, IServiceError)
	GenerateRefreshToken(user *entity.User) (string, IServiceError)
	GenerateAccountActivationToken(userId uint, email string) (string, IServiceError)
	GeneratePasswordResetToken(userId uint, email string) (string, IServiceError)
	ParseToken(token string) (*claims.Claims, error)
}
