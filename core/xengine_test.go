package core

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

import (
	"testing"
	"fmt"
	"github.com/XGoServer/config"
)

func TestConnectStr(t *testing.T) {
	fmt.Println( fmt.Sprintf(
		"%s:%s@/%s?charset=utf8","123","456","789"))
}

func TestCreateEngine(t *testing.T) {
	config.BindServerConfig("server.json","log.json")
	CreateDefaultMysqlEngine("mysql",config.ServerConfig.DbUser,config.ServerConfig.DbPw,config.ServerConfig.DbName)
}