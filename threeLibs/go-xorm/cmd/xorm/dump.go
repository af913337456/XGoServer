// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/XGoServer/threeLibs/go-xorm/core"
	"github.com/XGoServer/threeLibs/go-xorm/xorm"
)

var CmdDump = &Command{
	UsageLine: "dump driverName datasourceName",
	Short:     "dump database all table struct's and data to standard output",
	Long: `
dump database for sqlite3, mysql, postgres.

    driverName        Database driver name, now supported four: mysql mymysql sqlite3 postgres
    datasourceName    Database connection uri, for detail infomation please visit driver's project page
`,
}

func init() {
	CmdDump.Run = runDump
	CmdDump.Flags = map[string]bool{}
}

func runDump(cmd *Command, args []string) {
	if len(args) != 2 {
		fmt.Println("params error, please see xorm help dump")
		return
	}

	var err error
	engine, err = xorm.NewEngine(args[0], args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_UNKNOWN)

	err = engine.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = engine.DumpAll(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
}
