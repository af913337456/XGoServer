package util

import (
	"net/http"
	"encoding/json"
	l4g "github.com/XGoServer/threeLibs/alecthomas/log4go"
	"io/ioutil"
	"encoding/base32"
	"github.com/pborman/uuid"
	"bytes"
	"github.com/XGoServer/model"
	"net"
	"encoding/xml"
)

const isOpenDebug = true

func LogInfo(str string)  {
	if isOpenDebug {
		l4g.Info(str)
	}
}

func LogInterface(arg interface{})  {
	if isOpenDebug {
		l4g.Info(arg)
	}
}

func LogError(str string)  {
	if isOpenDebug {
		l4g.Error(str)
	}
}

func RenderJson(w http.ResponseWriter, o interface{}) {
	if b, err := json.Marshal(o); err != nil {
		w.Write([]byte(""))
	} else {
		w.Write(b)
	}
}

func BindJson(r *http.Request,params interface{}) map[string]interface{} {
	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody,params)

	if err != nil {
		return GetCommonErr("params error")
	}
	return nil
}

// 保持 html 格式输出 json
func JSONMarshalKeepHTML(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// 渲染成 xml
func RenderXml(w http.ResponseWriter, o interface{}) {
	if b, err := xml.Marshal(o); err != nil {
		w.Write([]byte(""))
	} else {
		w.Write(b)
	}
}

func GetCommonErr(info interface{}) map[string]interface{} {
	l4g.Error(info) // 保存一份到本地 log，方便日后排查问题
	d := map[string]interface{}{
		"errcode": 1,
		"errmsg" : info,
	}
	return d
}

func GetDefaultSuccess() map[string]interface{} {
	d := map[string]interface{}{
		"success": true,
	}
	return d
}

func GetCommonSuccess(info interface{}) map[string]interface{} {
	if info == nil {
		info = "null"
	}
	d := map[string]interface{}{
		"ret": "success",
		"msg":  info,
	}
	return d
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")
func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

func GetIpAddress(r *http.Request) string {
	address := r.Header.Get(model.HEADER_FORWARDED)

	if len(address) == 0 {
		address = r.Header.Get(model.HEADER_REAL_IP)
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return address
}