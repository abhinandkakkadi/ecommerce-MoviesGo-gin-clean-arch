package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

type authCustomClaimsUsers struct {
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenUsers(user domain.Users) (string,error) {

	claims := &authCustomClaimsUsers{
		Name: user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*48).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString([]byte("132457689"))

	if err != nil {
		return "",err
	}

	return tokenString,nil

}