// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQueryString(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type GetVar2 struct {
		Id      int64  `xorm:"autoincr pk"`
		Msg     string `xorm:"varchar(255)"`
		Age     int
		Money   float32
		Created time.Time `xorm:"created"`
	}

	assert.NoError(t, testEngine.Sync2(new(GetVar2)))

	var data = GetVar2{
		Msg:   "hi",
		Age:   28,
		Money: 1.5,
	}
	_, err := testEngine.InsertOne(data)
	assert.NoError(t, err)

	records, err := testEngine.QueryString("select * from get_var2")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(records))
	assert.Equal(t, 5, len(records[0]))
	assert.Equal(t, "1", records[0]["id"])
	assert.Equal(t, "hi", records[0]["msg"])
	assert.Equal(t, "28", records[0]["age"])
	assert.Equal(t, "1.5", records[0]["money"])
}

func TestQueryString2(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type GetVar3 struct {
		Id  int64 `xorm:"autoincr pk"`
		Msg bool  `xorm:"bit"`
	}

	assert.NoError(t, testEngine.Sync2(new(GetVar3)))

	var data = GetVar3{
		Msg: false,
	}
	_, err := testEngine.Insert(data)
	assert.NoError(t, err)

	records, err := testEngine.QueryString("select * from get_var3")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(records))
	assert.Equal(t, 2, len(records[0]))
	assert.Equal(t, "1", records[0]["id"])
	assert.True(t, "0" == records[0]["msg"] || "false" == records[0]["msg"])
}

func toString(i interface{}) string {
	switch i.(type) {
	case []byte:
		return string(i.([]byte))
	case string:
		return i.(string)
	}
	return fmt.Sprintf("%v", i)
}

func toInt64(i interface{}) int64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseInt(string(i.([]byte)), 10, 64)
		return n
	case int:
		return int64(i.(int))
	case int64:
		return i.(int64)
	}
	return 0
}

func toFloat64(i interface{}) float64 {
	switch i.(type) {
	case []byte:
		n, _ := strconv.ParseFloat(string(i.([]byte)), 64)
		return n
	case float64:
		return i.(float64)
	case float32:
		return float64(i.(float32))
	}
	return 0
}

func TestQueryInterface(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type GetVarInterface struct {
		Id      int64  `xorm:"autoincr pk"`
		Msg     string `xorm:"varchar(255)"`
		Age     int
		Money   float32
		Created time.Time `xorm:"created"`
	}

	assert.NoError(t, testEngine.Sync2(new(GetVarInterface)))

	var data = GetVarInterface{
		Msg:   "hi",
		Age:   28,
		Money: 1.5,
	}
	_, err := testEngine.InsertOne(data)
	assert.NoError(t, err)

	records, err := testEngine.QueryInterface("select * from get_var_interface")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(records))
	assert.Equal(t, 5, len(records[0]))
	assert.EqualValues(t, 1, toInt64(records[0]["id"]))
	assert.Equal(t, "hi", toString(records[0]["msg"]))
	assert.EqualValues(t, 28, toInt64(records[0]["age"]))
	assert.EqualValues(t, 1.5, toFloat64(records[0]["money"]))
}
