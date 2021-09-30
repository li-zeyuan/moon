package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretEnvKey = "moon_jwt_secret"

var SecretKey = os.Getenv(secretEnvKey)

type JwtClaims struct {
	Uid            int64
	Role           int
	StandardClaims jwt.StandardClaims
}

func (j JwtClaims) Valid() error {
	if len(strconv.FormatInt(j.Uid, 10)) != 18 {
		return errors.New("uid format error")
	}
	if !IsInArray(j.Role, AllRole) {
		return errors.New("role error")
	}
	if err := j.StandardClaims.Valid(); err != nil {
		return err
	}

	return nil
}

func GenerateToken(uid int64, role int, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Uid:  uid,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})

	return token.SignedString([]byte(SecretKey))
}
