package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

type MapClaims struct {
	Identidy string `json:"identidy"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

func MakeToken(identidy string, name string) (string, error) {
	UserMapClaims := &MapClaims{
		identidy,
		name,
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserMapClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString, err)
	return tokenString, nil
}


func ParseToken(tokenString string) (*MapClaims, error) {
	
	UserMapClaims := &MapClaims{}
	
	// 使用ParseWithClaims方法                   // 存放用户数据的变量
	token, err := jwt.ParseWithClaims(tokenString, UserMapClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	// 解析哪个结构体
	if claims, ok := token.Claims.(*MapClaims); ok && token.Valid {
		fmt.Println(claims)
		// 返回存放用户数据的变量 因为是结构体形式 所以 上面返回为*类型
		return UserMapClaims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
