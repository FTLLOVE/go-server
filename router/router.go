package router

import (
	"github.com/gin-gonic/gin"
	"go-server/controller"
)
func Gininit(router *gin.Engine)  {
	router = gin.Default()

	st:=router.Group("/open")
	st.GET("/hello", controller.GetTest)

}