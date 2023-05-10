package helper

import (
	"errors"
	"net/mail"
	"strings"
)

func ExtractBearerToken(Authheader string) (string, error) {
	if Authheader == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(Authheader, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func ValidateEmailPassword(email string, password string) bool {
	_, err := mail.ParseAddress(email)

	return len(password) > 5 && err == nil
}
