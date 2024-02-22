package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p []byte) string {
	psw, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(psw)
}

func ComparePassword(hashedPsw, psw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPsw, psw)

	return err == nil
}