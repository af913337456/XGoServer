package core

import (
	"net/http"
	"github.com/XGoServer/util"
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

func (x XHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	tokenStr := r.Header.Get(TokenAuth)

	c := &Context{}
	c.IpAddress = util.GetIpAddress(r)
	c.RoutePath = r.URL.Path

	if x.requireToken {
		tokenData,err := ParseToken(tokenStr)
		if err != nil {
			util.RenderJson(w,util.GetCommonErr(err.Error()))
			return
		}
		c.TokenData = *tokenData
		c.TokenStr  = tokenStr
		// util.RenderJson(w,c)
		x.handleFunc(c,w,r)
		return
	}
	// 不需要 token
	util.LogInterface(c)
	x.handleFunc(c,w,r)
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTZhYWEiLCJ1c2VyUm9sZSI6Im5vcm1hbCJ9.w-wBZWYYKgMQa50qZRc9qosiLaEtuK6t5b_tZjwhJ7A
func ApiRequestTokenHandler(f func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	x := XHandler{}
	x.handleFunc = f
	x.requireToken = true
	return x
}

func ApiNormalHandler(f func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	x := XHandler{}
	x.handleFunc = f
	x.requireToken = false
	return x
}
























