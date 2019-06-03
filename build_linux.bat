:: 编译脚本
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -a -v -o build/server serverMain.go