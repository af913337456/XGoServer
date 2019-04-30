package core

import (
	"net/http"
	"github.com/XGoServer/util"
)


const TokenAuth = "Authorization"
const SecretKey = "1234567890asdfghjklzxcvbnmqwert"

type Context struct {  // 可以自定义的基础上下文
	TokenData TokenData
	TokenStr  string  	`json:"tokenStr"`
	IpAddress string  	`json:"ipAddress"`
	RoutePath string  	`json:"routePath"`
}

type ReqContext struct {  // 公用请求上下文
	C *Context			  // 基础上下文
	W http.ResponseWriter // 输出
	R *http.Request       // 输入
}

type XHandler struct {
	handleFunc   func(*ReqContext)
	requireToken bool
}

// 自定义封装实现 handler 接口
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
		x.handleFunc(&ReqContext{c,w,r})
		return
	}
	// 不需要 token
	util.LogInterface(c)
	x.handleFunc(&ReqContext{c,w,r})
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTZhYWEiLCJ1c2VyUm9sZSI6Im5vcm1hbCJ9.w-wBZWYYKgMQa50qZRc9qosiLaEtuK6t5b_tZjwhJ7A
func ApiRequestTokenHandler(f func(*ReqContext)) http.Handler {
	x := XHandler{}
	x.handleFunc = f
	x.requireToken = true
	return x
}

func ApiNormalHandler(f func(*ReqContext)) http.Handler {
	x := XHandler{}
	x.handleFunc = f
	x.requireToken = false
	return x
}
























