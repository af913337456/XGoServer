package core

import (
	"net/http"
	"github.com/XGoServer/util"
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

