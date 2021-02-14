package claims

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Id       uint
	HashedId string
	Email    string
	Username string
	jwt.StandardClaims
}
