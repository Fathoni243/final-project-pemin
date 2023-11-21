package util

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidationPassword(password string) error {
	var (
		numeric, lowerCase, upperCase, specialCharacter bool
	)

	numeric = regexp.MustCompile(`\d`).MatchString(password)
	lowerCase = regexp.MustCompile(`[a-z]`).MatchString(password)
	upperCase = regexp.MustCompile(`[A-Z]`).MatchString(password)
	specialCharacter = strings.ContainsAny(password, "!@#$%^&*()_+-=/.,:;'`?{}[|]")

	if len(password) < 6 {
		return errors.New("Too short password")
	}

	if !numeric {
		return errors.New("password need numeric character")
	}

	if !lowerCase {
		return errors.New("password need lower case character")
	}

	if !upperCase {
		return errors.New("password need upper case character")
	}

	if !specialCharacter {
		return errors.New("password need spesial character ex !@#$%^&*()_+-=/.,:;'`?{}[|]")
	}
	
	return nil
}
