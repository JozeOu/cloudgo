package main

import (
	"os"

	"github.com/github-user/web/service"

	flag "github.com/spf13/pflag"
)

const (
	// 默认端口号：8080
	PORT string = "8080"
)

func main() {
	// 从环境变量中获取端口号
	port := os.Getenv("PORT")
	// 若端口号不存在，则将端口号设置为默认端口号
	if len(port) == 0 {
		port = PORT
	}
	//对命令参数进行绑定
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// 创建服务器端
	server := service.NewServer()
	// 设置监听端口（使用negroni中的Run函数）
	server.Run(":" + port)
}
