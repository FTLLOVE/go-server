package gin_server

import (
	//后续增加日志
	_"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"go-server/middleware"
	"io"
	"os"
)
func StartHttpServer(listen string)  {

	//控制台颜色
	gin.ForceConsoleColor()
	//日志
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	router:=gin.Default()//暂时使用default

	//增加中间件
	router.Use(middleware.GetUseTime)

	st:=router.Group("/open")
	st.GET("/hello", GetTest)
	st.POST("/register",RegisterUser)

	//异步消费运行
	RunConsumer()
	//start
	if err := router.Run(listen); err != nil {
		println("Error when running server. " + err.Error())
	}

}