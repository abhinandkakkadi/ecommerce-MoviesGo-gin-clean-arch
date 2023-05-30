package helper

import (
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {

	claims := &authCustomClaimsAdmin{
		Name:  admin.Name,
		Email: admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("12345678"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
