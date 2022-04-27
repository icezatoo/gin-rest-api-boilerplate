package bcrypt

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return string(bytes)
}

func ComparePassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
