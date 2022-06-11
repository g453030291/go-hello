# go-hello

# 环境配置
`go env`
`go env -W GOPROXY=https://goproxy.cn,direct`
`go env -W GO111MODULE=on`

# 初始化项目
`go mod init go-hello`
`go mod tidy`

# 编译执行
`go build .\hello.go`
`.\hello.exe`
`go run .\hello.go`

# go语言常用的api及小示例
1.file-demo 文件读写(参考https://developer.51cto.com/article/697112.html)