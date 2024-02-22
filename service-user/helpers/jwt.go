package helpers

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var jwt_secret string = "this_is_jwt_secret_shhh"

func SignToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	tokenString, err := token.SignedString([]byte(jwt_secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(jwt_secret), nil
	})

	if err != nil {
		return nil, err
	}


	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("JWT ERROR")
	}

	return claims, nil
}