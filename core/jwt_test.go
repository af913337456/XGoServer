package core

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/16

*/

import (
	"testing"
	"github.com/XGoServer/threeLibs/dgrijalva/jwt-go"
	"fmt"
	"github.com/XGoServer/threeLibs/dgrijalva/jwt-go/request"
	"net/http"
)

const SecretKey = "1234567890asdfghjklzxcvbnmqwert"

func TestJWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	dataMap := make(jwt.MapClaims)
	dataMap["userId"]   = "123456aaa"
	dataMap["userRole"] = "normal"
	token.Claims = dataMap

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(tokenString)
	}
	parseToken(tokenString)
}

func parseToken(tokenStr string) {
	h := http.Header{"Authorization":[]string{tokenStr}}
	r := &http.Request{
		Header:h,
	}
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			fmt.Println("correct")
			fmt.Println(token.Claims)
			return
		}
	}
	fmt.Println(err.Error())
}
