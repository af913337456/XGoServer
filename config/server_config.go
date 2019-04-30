package config

import (
	io "io/ioutil"
	json "encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type configer struct{

}

func NewConfiger () *configer {
	return &configer{}
}

func (s *configer) Load (filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Load config json failed ===> filename : %s %s",filename,err.Error()))
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil{
		panic(fmt.Sprintf("read json failed ===> filename : %s %s",filename,err.Error()))
	}
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

func BindServerConfig(serverConfigFileName,logConfigFileName string) {
	configer := NewConfiger()
	/** 传入的 结构体 要和 json 的格式对上,否则返回是 null */
	configer.Load(FindConfigFile(serverConfigFileName),  &ServerConfig)
	configer.Load(FindConfigFile(logConfigFileName), &LogConfig)
	jsonBytes,_ := json.Marshal(&ServerConfig)
	fmt.Println(string(jsonBytes))
	jsonBytes,_ = json.Marshal(&LogConfig)
	fmt.Println(string(jsonBytes))
}

func FindConfigFile(fileName string) string {
	if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat("../../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../../config/" + fileName)
	}else if _, err := os.Stat("../../../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../../../config/" + fileName)
	}else if _, err := os.Stat("../../../../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../../../../config/" + fileName)
	} else if _, err := os.Stat("config/"+fileName); err == nil {
		fileName, _ = filepath.Abs("config/"+fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}
	return fileName
}