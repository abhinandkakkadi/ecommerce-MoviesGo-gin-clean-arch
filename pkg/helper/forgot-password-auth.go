package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ExtractUserIDFromTokenForgotPassword(tokenString string) (int, string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsUsers{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte("reset"), nil // Replace with your actual secret key
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*authCustomClaimsUsers)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	return claims.Id, claims.Email, nil

}
