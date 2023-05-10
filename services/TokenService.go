package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	ID    int       `json:"id"`
	Email string    `json:"Email"`
	Exp   time.Time `json:"exp"`
	jwt.StandardClaims
}

func CreateToken(id int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		ID:             id,
		Email:          email,
		Exp:            time.Now().AddDate(0, 0, 7),
		StandardClaims: jwt.StandardClaims{},
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return signedToken, err

}

func VerifyToken(token string) (*TokenClaims, bool) {
	parsed, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	v, ok := parsed.Claims.(*TokenClaims)
	if ok {
		if v.Exp.Before(time.Now()) {
			return nil, false
		}
		return v, true
	}
	return nil, false
}
