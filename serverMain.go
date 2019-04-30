package main
/**
  * 作者：林冠宏
  *
  * author: LinGuanHong
  *
  * My GitHub : https://github.com/af913337456/
  *
  * My Blog   : http://www.cnblogs.com/linguanh/
  *
  * */

/**
	模板 main.go
*/

import (
	"github.com/XGoServer/threeLibs/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/XGoServer/core"
	"github.com/XGoServer/model"
	"github.com/XGoServer/util"
	"github.com/XGoServer/encrypt"
)

func setRouter() *mux.Router {
	router := new (mux.Router)
	router.HandleFunc("/",test).Methods("POST")
	router.HandleFunc("/test",test).Methods("GET")
	router.HandleFunc("/test2",test2).Methods("GET")
	router.HandleFunc("/test3",test3).Methods("GET")

	router.Handle("/fuck",core.ApiNormalHandler(getToken)).Methods("GET")
	router.Handle("/check",core.ApiRequestTokenHandler(handleToken)).Methods("GET")

	router.HandleFunc("/enc",encOutput).Methods("GET")
	router.HandleFunc("/dec",decOutput).Methods("POST")
	/** 在下面添加你的回调方法 */
	/** add your func below */
	return router
}

func main()  {
	core.HttpListen(setRouter())
}

func test(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"======= hello world! =======")
}

func test2(w http.ResponseWriter,r *http.Request)  {
	// 非常简单的例子
	// 操作放在内部 , 可以使用 request 来获取自己的参数，再直接组织输出
	core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
		m :=  map[string]interface{}{}
		m["fuck"] = "fuck"
		return m
	})
}

// 获取 token
func getToken(req *core.ReqContext)  {

	core.HandlerMapWithOutputJson(req.W, func() map[string]interface{} {
		tokenStr,err := core.BuildDefaultToken(func(tokenData *core.TokenData) {
			tokenData.UserId = "123456"
			tokenData.Roles  = "normal"
		})
		if err != nil {
			return util.GetCommonErr(err.Error())
		}
		return util.GetCommonSuccess(tokenStr)
	})

}

// 解析输出 token
func handleToken(req *core.ReqContext)  {
	util.RenderJson(req.W,req.C.TokenData)
}

// 加密输出
func encOutput(w http.ResponseWriter, r *http.Request)  {
	core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
		d := "狗年平安"
		aes := encrypt.DefaultAES{}
		return util.GetCommonSuccess(aes.AesEncryptStr(d))
	})
}

//解密输出
func decOutput(w http.ResponseWriter, r *http.Request)  {
	core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
		d := r.PostFormValue("enc")
		if d == "" {
			return util.GetCommonErr("empty")
		}
		aes := encrypt.DefaultAES{}
		util.LogInfo("d ====> "+d)
		return util.GetCommonSuccess(aes.AesDecryptStr(d))
	})
}

func test3(w http.ResponseWriter,r *http.Request)  {
	// 加上 xorm 来演示
	core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
		// 插入一条评论
		item := &model.Comment{
			Id		:util.NewId(),
			UserId	:"123456",
			Name	:"LinGuanHong",
			Content	:"hello word",
		}
		affect,_ := core.Engine.Insert(item)
		m :=  map[string]interface{}{}
		if affect > 0 {
			m["ret"] = "insert success"
			// 获取所有评论输出
			comments := make([]model.Comment, 0)
			core.Engine.Find(&comments)
			m["msg"] = comments
		}else{
			m["ret"] = "insert failed"
		}
		return m
	})
}


















