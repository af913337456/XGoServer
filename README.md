> 作者：林冠宏 / 指尖下的幽灵

> 掘金：https://juejin.im/user/587f0dfe128fe100570ce2d8

> 博客：http://www.cnblogs.com/linguanh/

---

<strong>一个基础性、模块完整且安全可靠的服务端框架</strong>

#### 你可以使用它
* 简单快速搭建自己的服务端
* 其他高级模块拓展等

#### 具备的

* Token模块，``jwt``
* 加解密模块，``cipher-AES``，可自行拓展其他
* 日志模块，``alecthomas/log4go``
* 路由模块，``gorilla/mux``
* 硬存储 / 软存储 采用 ``xorm`` 框架
* 服务端通用的输出数据结构的整合，例如 json
* 各模块对应的单元测试例子

##### 如果你想直接输出一条 json 给客户端，这样子

```golang
func main()  {
    router := new (mux.Router)
    router.HandleFunc("/",test2).Methods("GET")
    core.HttpListen(router)
}
func test2(w http.ResponseWriter,r *http.Request)  {
    // 非常简单的例子, 操作放在内部 , 可以使用 request 来获取自己的参数，再直接组织输出
    core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
    	m :=  map[string]interface{}{}
    	m["msg"] = "blow me a kiss"
    	return m
    })
}
// 结果 ： {"msg":"blow me a kiss"}
```

##### 令牌机制

```golang

// core.ApiNormalHandler 不要求在请求头中传递 Token
router.Handle("/fuck",core.ApiNormalHandler(getToken)).Methods("GET")

// core.ApiRequestTokenHandler 要求在请求头中带上 Token
router.Handle("/check",core.ApiRequestTokenHandler(handleToken)).Methods("GET")

```

##### 与数据库交互

```go
func test3(w http.ResponseWriter,r *http.Request)  {
	core.HandlerMapWithOutputJson(w, func() map[string]interface{} {
		// 插入一条评论
		item := &model.Comment{
			Id	:util.NewId(),         // 评论 id
			UserId	:"123456",             // 评论人 id
			Name	:"LinGuanHong",        // 评论人名称
			Content	:"hello word",         // 评论内容
		}
		affect,_ := core.Engine.Insert(item)  // 执行插入，传入 struct 引用
		m :=  map[string]interface{}{}
		if affect > 0 {
			m["ret"] = "insert success"
			comments := make([]model.Comment, 0)
			core.Engine.Find(&comments)   // select 出来，获取所有评论输出
			m["msg"] = comments
		}else{
			m["ret"] = "insert failed"
		}
		return m
	})
}

输出的结果是：
{
  "msg": [
    {
      "id": "1kubpgh9pprrucy11e456fyytw",
      "UserId": "123456",
      "name": "LinGuanHong",
      "content": "hello word"
    }
  ],
  "ret": "insert success"
}

```

### 使用流程
目录如下

```go
---- config
---- core
---- model
---- threeLibs
---- util
---- server.go
```

1 在 ``config`` 放置配置文件

* ``服务端配置 json 文件`` -- server.json，
* ``日志配置文件`` -- log.json 例如下面的，他们都会在``运行程序后会自动解析和读取``

2 ``threeLibs`` 目录放置了依赖的第三方库，例如 xorm，不需要你再去 go get

3 ``model`` 放置数据实体 struct

```go
{
  "Host": "127.0.0.1",
  "Port": ":8884",
  "FilePort":":8885",
  "DbName":"lgh",
  "DbUser":"root",
  "DbPw":"123456",
  "DbPort":"3306"
}
```

```go
{
  "EnableConsole": true,
  "ConsoleLevel": "DEBUG",
  "EnableFile": true,
  "FileLevel": "INFO",
  "FileFormat": "",
  "FileLocation": ""
}
```


### 从一个最基础的例子开始：

```golang
func main()  {
    router := new (mux.Router)
    router.HandleFunc("/",test).Methods("GET")
    /** 在下面添加你的路由 */
    /** add your routine func below */
    core.HttpListen(router)  // 简单的 http 监听，当然也提供了 https
}
func test(w http.ResponseWriter,r *http.Request)  {
    fmt.Fprintf(w,"======= hello world! =======")
}

```

```golang
// http 监听
func HttpListen(router *mux.Router)  {
	SimpleInit()  // 此处自动初始化 ---------- ①
	url := config.ServerConfig.Host+config.ServerConfig.Port
	util.LogInfo("服务启动于 : "+url)
	err := http.ListenAndServe(url,router)
	if err !=nil {
		util.LogInfo("http error ===> : "+err.Error())
		return
	}
}
```

```golang
// 绑定配置 json 的信息 以及 初始化 xorm mysql数据库引擎
func SimpleInit() {
	config.BindServerConfig("server.json","log.json")
	fmt.Println("BindServerConfig ==================>","success")
	config.ConfigureLog(&config.LogConfig)
	CreateDefaultMysqlEngine(
		"mysql",
		config.ServerConfig.DbUser,
		config.ServerConfig.DbPw,
		config.ServerConfig.DbName)
}
```

### 服务端通用的输出数据结构的整合函数组

```go
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

// 还有
// HandlerStructWithOutputXML  // XML 的输出格式
// HandlerStructWithOutputString

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
```

就介绍这么多了

