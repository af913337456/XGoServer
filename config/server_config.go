package config

import (
	io "io/ioutil"
	json "encoding/json"
	"fmt"
)

type configer struct{

}

func NewConfiger () *configer {
	return &configer{}
}

func (s *configer) Load (filename string, v interface{}) bool {
	data, err := io.ReadFile(filename)
	if err != nil {
		fmt.Println("Load config json failed ===> filename : "+filename+" -- "+err.Error())
		return false
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil{
		fmt.Println("read json failed ===> filename : "+filename+" -- "+err.Error())
		return false
	}
	return true
}

type ServerConfigStruct struct{
	Host 	 string `json:"host"`
	Port     string `json:"port"`
	FilePort string `json:"filePort"`
	DbName 	 string `json:"dbName"`
	DbUser 	 string `json:"dbUser"`
	DbPw   	 string `json:"dbPw"`
	DbPort 	 string `json:"dbPort"`
}

type LogConfigStruct struct {
	EnableConsole bool   `json:"EnableConsole"`
	ConsoleLevel  string `json:"ConsoleLevel"`
	EnableFile    bool   `json:"EnableFile"`
	FileLevel     string `json:"FileLevel"`
	FileFormat    string `json:"FileFormat"`
	FileLocation  string `json:"FileLocation"`
}

var ServerConfig  ServerConfigStruct
var LogConfig LogConfigStruct

func BindServerConfig() bool {
	configer := NewConfiger()
	/** 传入的 结构体 要和 json 的格式对上,否则返回是 null */
	isDbSuccess  := configer.Load("config/server.json",  &ServerConfig)
	isLogSuccess := configer.Load("config/log.json", &LogConfig)
	if !isDbSuccess || !isLogSuccess {
		return false
	}
	jsonBytes,_ := json.Marshal(&ServerConfig)
	fmt.Println(string(jsonBytes))
	jsonBytes,_ = json.Marshal(&LogConfig)
	fmt.Println(string(jsonBytes))
	return true
}
