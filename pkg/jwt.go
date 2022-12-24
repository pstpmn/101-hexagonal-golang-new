package pkg

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type IJwt interface {
	Generate(data map[string]interface{}, key string, exp time.Time) (string, error)
	Extract(token string, key string) (map[string]string, error)
}

type j struct {
}

func NewJsonWebToken() IJwt {
	return &j{}
}

func (j j) Generate(data map[string]interface{}, key string, exp time.Time) (string, error) {
	data["exp"] = exp.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	t, err := token.SignedString([]byte(key))
	return t, err
}

func (j j) Extract(clientToken string, key string) (map[string]string, error) {
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims)
		return nil, nil
	} else {
		return nil, err
	}
}
