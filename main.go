package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/router"
)

const port = 9527

func main(){
	route:= gin.New()
	router.Gininit(route)
	if err := route.Run(fmt.Sprintf(":%d", port)); err != nil {
		println("Error when running server. " + err.Error())
	}
}