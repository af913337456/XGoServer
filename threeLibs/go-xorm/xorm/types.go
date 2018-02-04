package xorm

import (
	"reflect"

	"github.com/XGoServer/threeLibs/go-xorm/core"
)

var (
	ptrPkType = reflect.TypeOf(&core.PK{})
	pkType    = reflect.TypeOf(core.PK{})
)
