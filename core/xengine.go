package core

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

import (
	"github.com/XGoServer/threeLibs/go-xorm/xorm"
	"github.com/XGoServer/threeLibs/go-xorm/core"
	l4g "github.com/XGoServer/threeLibs/alecthomas/log4go"
	_ "github.com/XGoServer/threeLibs/go-sql-driver/mysql"
	"fmt"
)

const DbPrefix     = "lgh_"
const DbEngineType = "InnoDB"
const DbChartSet   = "utf8"
var Engine *xorm.Engine

// 可以自行根据 xorm 支持的类型修改数据库
// 默认的采用 mysql
func CreateDefaultMysqlEngine(dbType,dbUser,dbPw,dbName string) bool {
	var err error
	connectStr := fmt.Sprintf("%s:%s@/%s?charset=utf8",dbUser,dbPw,dbName)
	Engine, err = xorm.NewEngine(dbType,connectStr)
	if err != nil {
		l4g.Error("CreateDefaultMysqlEngine failed ===> %s",err.Error())
		return false
	}
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_DEBUG)
	Engine.SetMaxIdleConns(6)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, DbPrefix)
	Engine.SetTableMapper(tbMapper)
	Engine.StoreEngine(DbEngineType)
	Engine.Charset(DbChartSet)
	err = Engine.Ping()
	if err != nil {
		l4g.Info("create default XEngine failed : ===> %s",err.Error())
		return false
	}
	l4g.Info("create default XEngine success : ===> %s",connectStr)
	return true
}

func CreateTables(beans ...interface{}) bool {
	if Engine == nil {
		l4g.Error("engine is null,create tables failed!")
		return false
	}
	err := Engine.CreateTables(beans...)
	if err != nil {
		l4g.Error("CreateTables failed ",err.Error())
		return false
	}
	return true
}
