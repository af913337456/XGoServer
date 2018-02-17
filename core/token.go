package core

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/17

*/

import (
	"github.com/XGoServer/threeLibs/dgrijalva/jwt-go"
	"time"
	"fmt"
)

const TokenExpiresTime = 60//60*24*60*100*2

type TokenError struct {
	ErrorInfo string
}

func (t TokenError) Error() string {
	return t.ErrorInfo
}

type TokenData struct {
	UserId string            `json:"user_id"`
	Roles  string            `json:"roles"`
	Props  map[string]string `json:"props"`
	jwt.StandardClaims
}

func BuildDefaultToken(f func(tokenData *TokenData)) (string,error) {
	timeNow := time.Now().Unix()
	tokenData := &TokenData{
		StandardClaims:jwt.StandardClaims{
			IssuedAt :timeNow,
			ExpiresAt:timeNow+TokenExpiresTime,
		},
	}
	f(tokenData)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,tokenData)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "",err
	}
	return tokenString,nil
}

func ParseToken(tokenStr string) (*TokenData,error) {
	if len(tokenStr) == 0 {
		// error 没有 token
		return nil,TokenError{ErrorInfo:"illegal token"}
	}
	// parse token
	token, err := jwt.ParseWithClaims(tokenStr, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token method")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil,TokenError{ErrorInfo:"parse token err: "+err.Error()}
	}

	tokenData := token.Claims.(*TokenData)

	if !token.Valid || tokenData.ExpiresAt == 0 {
		return nil,TokenError{ErrorInfo:"invalid token"}
	}
	// 不需要自己再去判断过期，内部判断了 claims.go --> valid()
	// now := time.Now().Unix()
	// if tokenData.ExpiresAt < now {
	//	 util.RenderJson(w,util.GetCommonErr("expires token"))
	//	 return
	// }
	return tokenData,nil
}




























