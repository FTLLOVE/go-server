package router

import (
	"github.com/gin-gonic/gin"
	"go-server/controller"
)
func Gininit(router *gin.Engine)  {

	st:=router.Group("/open")
	st.GET("/hello", controller.GetTest)

}