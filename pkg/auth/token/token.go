package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type UserClaims struct {
	UserID int64
	jwt.RegisteredClaims
}

func New(userID int64) (string, error) {
	c := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &c)
	k := []byte(os.Getenv("TOKEN_KEY"))
	if len(k) == 0 {
		return "", fmt.Errorf("%w: key is empty", jwt.ErrInvalidKey)
	}
	ss, err := token.SignedString(k)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func Validate(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			k := os.Getenv("TOKEN_KEY")
			if len(k) == 0 {
				return nil, fmt.Errorf("%w: key is empty", jwt.ErrInvalidKey)
			}
			return []byte(k), nil
		})
	if err != nil {
		return err
	}
	if c, ok := token.Claims.(*UserClaims); ok {
		if c.ExpiresAt.After(time.Now()) {
			return jwt.ErrTokenExpired
		}
	}
	return nil
}
