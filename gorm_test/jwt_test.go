package gormtest

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func TestToken(t *testing.T) {
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	fmt.Println(tokenString, err)
}

func TestParseToken(t *testing.T) {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ._NaFhGu8tCCgBKksGBA6ADwRdKx3e9GES_KyF4A5phE"
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
}
