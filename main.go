package main

import (
	"go-server/gin_server"
)
 var listen string = "9527"

func main(){
	//启动服务
	gin_server.StartHttpServer(listen)

}