package core

import (
	"fmt"
	"github.com/XGoServer/config"
	"github.com/XGoServer/threeLibs/gorilla/mux"
	"crypto/x509"
	"io/ioutil"
	"crypto/tls"
	"net/http"
	"github.com/XGoServer/util"
)

func SimpleInit() bool {
	if config.BindServerConfig() {
		fmt.Println("BindServerConfig ==================> success")
		config.ConfigureLog(&config.LogConfig)
		CreateDefaultMysqlEngine(
			"mysql",
			config.ServerConfig.DbUser,
			config.ServerConfig.DbPw,

			config.ServerConfig.DbName)
		return true
	}else{
		fmt.Println("BindServerConfig ===> failed")
		return false
	}
}

// http 监听
func HttpListen(router *mux.Router)  {
	SimpleInit()
	url := config.ServerConfig.Host+config.ServerConfig.Port
	util.LogInfo("服务启动于 : "+url)
	err := http.ListenAndServe(url,router)
	if err !=nil {
		util.LogInfo("http 监听错误 : "+err.Error())
		return
	}
}

// https 监听
// ca.crt / server.crt / server.key 共三个文件
func HttpsListen(router *mux.Router,caCrt, serveCrt,serverKey string)  {

	SimpleInit()

	basePath := "" // /home/lgh/
	pool := x509.NewCertPool()
	caCertPath := basePath+caCrt // ca.crt

	caCrtBytes, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		util.LogInfo("https Read ca File err : "+err.Error())
		return
	}
	pool.AppendCertsFromPEM(caCrtBytes)
	s := &http.Server{
		Addr:    config.ServerConfig.Host+config.ServerConfig.Port, // :8888
		Handler: router,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert, /** 开启双向验证 */
		},
	}
	s.ListenAndServeTLS(basePath+serveCrt,basePath+serverKey)
}
