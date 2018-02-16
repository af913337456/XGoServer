package core

import (
	"net/http"
	"github.com/XGoServer/util"
	"github.com/XGoServer/threeLibs/dgrijalva/jwt-go"
	"fmt"
)

type Context struct {
	TokenData TokenData
	TokenStr  string  	`json:"tokenStr"`
	IpAddress string  	`json:"ipAddress"`
	RoutePath string  	`json:"routePath"`
}


const TokenAuth = "Authorization"
const SecretKey = "1234567890asdfghjklzxcvbnmqwert"

type XHandler struct {
	handleFunc     func(*Context, http.ResponseWriter, *http.Request)
	requireToken bool
}

type TokenData struct {
	UserId string            `json:"user_id"`
	Roles  string            `json:"roles"`
	Props  map[string]string `json:"props"`
	jwt.StandardClaims
}

func (x XHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	tokenStr := r.Header.Get(TokenAuth)

	c := &Context{}
	c.IpAddress = util.GetIpAddress(r)
	c.RoutePath = r.URL.Path

	if x.requireToken {
		if len(tokenStr) == 0 {
			// error 没有 token
			util.RenderJson(w,util.GetCommonErr("illegal token"))
			return
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
			util.RenderJson(w,util.GetCommonErr("parse token err: "+err.Error()))
			return
		}
		tokenData := token.Claims.(*TokenData)
		if !token.Valid {
			util.RenderJson(w,util.GetCommonErr("invalid token"))
			return
		}
		c.TokenData = *tokenData
		c.TokenStr  = tokenStr
		c.TokenData.ExpiresAt = 1234567
		util.RenderJson(w,c)
		x.handleFunc(c,w,r)
		return
	}
	// 不需要 token
	x.handleFunc(c,w,r)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTZhYWEiLCJ1c2VyUm9sZSI6Im5vcm1hbCJ9.w-wBZWYYKgMQa50qZRc9qosiLaEtuK6t5b_tZjwhJ7A
func MyHandlerFunc(f func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	x := XHandler{}
	x.handleFunc = f
	x.requireToken = true
	return x
}






















