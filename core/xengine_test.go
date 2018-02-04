package core

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
	config.BindServerConfig()
	CreateDefaultMysqlEngine("mysql",config.ServerConfig.DbUser,config.ServerConfig.DbPw,config.ServerConfig.DbName)
}