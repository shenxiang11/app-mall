package jwt_helper

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type DecodeToken struct {
	Id  uint `json:"id"`
	iat int  `json:"iat"`
	iss int  `json:"iss"`
}

func GenerateToken(tkn *jwt.Token, secret string) (string, error) {
	token, err := tkn.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(token string, secret string) (*DecodeToken, error) {
	decoded, err := jwt.Parse(token, func(tkn *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return &DecodeToken{}, err
	}

	if !decoded.Valid {
		return &DecodeToken{}, errors.New("VerifyToken invalid")
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodeToken
	jsonString, err := json.Marshal(decodedClaims)
	if err != nil {
		return &DecodeToken{}, errors.New("VerifyToken marshal failed")
	}
	err = json.Unmarshal(jsonString, &decodedToken)
	if err != nil {
		return &DecodeToken{}, errors.New("VerifyToken unmarshal failed")
	}

	return &decodedToken, nil
}
