package token

import (
	"errors"
	"log"
	"time"

	setting "github.com/Babatunde50/lms/pkg/settings"
	"github.com/golang-jwt/jwt/v4"
)

var ErrInvalidToken = errors.New("token is invalid")

type JWTMaker struct {
	secretKey string
}

var JWTTokenMaker *JWTMaker

func SetupJWTMaker() {
	if len(setting.TokenSetting.SecretKey) < setting.TokenSetting.MinSecretKeySize {
		log.Fatal("Invalid secret key")
	}

	JWTTokenMaker = &JWTMaker{secretKey: setting.TokenSetting.SecretKey}

}

func (maker *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)

	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
