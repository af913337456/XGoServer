package core

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/10

*/

import (
	"net/http"
	"github.com/XGoServer/util"
	"fmt"
)


type FinalResult struct {
	Data interface{}
}

type RetChannel chan FinalResult

func HandlerStruct(handle func() interface{}) *interface{} {
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	return ret.Data.(*interface{})
}

func HandlerMap(handle func() map[string]interface{}) *map[string]interface{} {
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	return ret.Data.(*map[string]interface{})
}

func HandlerStructWithOutputJson(w http.ResponseWriter,handle func() interface{})  {
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	mapRet := ret.Data.(*interface{})
	util.RenderJson(w,mapRet)
}

// 输出 xml 结果
func HandlerStructWithOutputXML(w http.ResponseWriter,handle func() interface{}){
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	mapRet := ret.Data.(*interface{})
	util.RenderXml(w,mapRet)
}

func HandlerStructWithOutputString(w http.ResponseWriter,handle func() string) {
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	dataStr := (ret.Data).(*string)
	fmt.Fprint(w,*dataStr)
}

func HandlerMapWithOutputJson(w http.ResponseWriter,handle func() map[string]interface{}){
	RetChannel := make(RetChannel, 1)
	go func() {
		result := FinalResult{}
		data := handle()
		result.Data = &data
		RetChannel <- result
		close(RetChannel)
	}()
	ret := <-RetChannel
	mapRet := ret.Data.(*map[string]interface{})
	util.RenderJson(w,mapRet)
}

