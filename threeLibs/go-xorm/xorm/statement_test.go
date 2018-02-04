// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"reflect"
	"strings"
	"testing"

	"github.com/XGoServer/threeLibs/go-xorm/core"
)

var colStrTests = []struct {
	omitColumn        string
	onlyToDBColumnNdx int
	expected          string
}{
	{"", -1, "`ID`, `IsDeleted`, `Caption`, `Code1`, `Code2`, `Code3`, `ParentID`, `Latitude`, `Longitude`"},
	{"Code2", -1, "`ID`, `IsDeleted`, `Caption`, `Code1`, `Code3`, `ParentID`, `Latitude`, `Longitude`"},
	{"", 1, "`ID`, `Caption`, `Code1`, `Code2`, `Code3`, `ParentID`, `Latitude`, `Longitude`"},
	{"Code3", 1, "`ID`, `Caption`, `Code1`, `Code2`, `ParentID`, `Latitude`, `Longitude`"},
	{"Longitude", 1, "`ID`, `Caption`, `Code1`, `Code2`, `Code3`, `ParentID`, `Latitude`"},
	{"", 8, "`ID`, `IsDeleted`, `Caption`, `Code1`, `Code2`, `Code3`, `ParentID`, `Latitude`"},
}

func TestColumnsStringGeneration(t *testing.T) {
	if dbType == "postgres" || dbType == "mssql" {
		return
	}

	var statement *Statement

	for ndx, testCase := range colStrTests {
		statement = createTestStatement()

		if testCase.omitColumn != "" {
			statement.Omit(testCase.omitColumn)
		}

		columns := statement.RefTable.Columns()
		if testCase.onlyToDBColumnNdx >= 0 {
			columns[testCase.onlyToDBColumnNdx].MapType = core.ONLYTODB
		}

		actual := statement.genColumnStr()

		if actual != testCase.expected {
			t.Errorf("[test #%d] Unexpected columns string:\nwant:\t%s\nhave:\t%s", ndx, testCase.expected, actual)
		}
		if testCase.onlyToDBColumnNdx >= 0 {
			columns[testCase.onlyToDBColumnNdx].MapType = core.TWOSIDES
		}
	}
}

func BenchmarkColumnsStringGeneration(b *testing.B) {
	b.StopTimer()

	statement := createTestStatement()

	testCase := colStrTests[0]

	if testCase.omitColumn != "" {
		statement.Omit(testCase.omitColumn) // !nemec784! Column must be skipped
	}

	if testCase.onlyToDBColumnNdx >= 0 {
		columns := statement.RefTable.Columns()
		columns[testCase.onlyToDBColumnNdx].MapType = core.ONLYTODB // !nemec784! Column must be skipped
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		actual := statement.genColumnStr()

		if actual != testCase.expected {
			b.Errorf("Unexpected columns string:\nwant:\t%s\nhave:\t%s", testCase.expected, actual)
		}
	}
}

func BenchmarkGetFlagForColumnWithICKey_ContainsKey(b *testing.B) {

	b.StopTimer()

	mapCols := make(map[string]bool)
	cols := []*core.Column{
		{Name: `ID`},
		{Name: `IsDeleted`},
		{Name: `Caption`},
		{Name: `Code1`},
		{Name: `Code2`},
		{Name: `Code3`},
		{Name: `ParentID`},
		{Name: `Latitude`},
		{Name: `Longitude`},
	}

	for _, col := range cols {
		mapCols[strings.ToLower(col.Name)] = true
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {

		for _, col := range cols {

			if _, ok := getFlagForColumn(mapCols, col); !ok {
				b.Fatal("Unexpected result")
			}
		}
	}
}

func BenchmarkGetFlagForColumnWithICKey_EmptyMap(b *testing.B) {

	b.StopTimer()

	mapCols := make(map[string]bool)
	cols := []*core.Column{
		{Name: `ID`},
		{Name: `IsDeleted`},
		{Name: `Caption`},
		{Name: `Code1`},
		{Name: `Code2`},
		{Name: `Code3`},
		{Name: `ParentID`},
		{Name: `Latitude`},
		{Name: `Longitude`},
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {

		for _, col := range cols {

			if _, ok := getFlagForColumn(mapCols, col); ok {
				b.Fatal("Unexpected result")
			}
		}
	}
}

type TestType struct {
	ID        int64   `xorm:"ID PK"`
	IsDeleted bool    `xorm:"IsDeleted"`
	Caption   string  `xorm:"Caption"`
	Code1     string  `xorm:"Code1"`
	Code2     string  `xorm:"Code2"`
	Code3     string  `xorm:"Code3"`
	ParentID  int64   `xorm:"ParentID"`
	Latitude  float64 `xorm:"Latitude"`
	Longitude float64 `xorm:"Longitude"`
}

func (TestType) TableName() string {
	return "TestTable"
}

func createTestStatement() *Statement {
	if engine, ok := testEngine.(*Engine); ok {
		statement := &Statement{}
		statement.Init()
		statement.Engine = engine
		statement.setRefValue(reflect.ValueOf(TestType{}))

		return statement
	} else if eg, ok := testEngine.(*EngineGroup); ok {
		statement := &Statement{}
		statement.Init()
		statement.Engine = eg.Engine
		statement.setRefValue(reflect.ValueOf(TestType{}))

		return statement
	}
	return nil
}
