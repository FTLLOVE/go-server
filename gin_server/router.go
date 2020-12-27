package gin_server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//后续增加日志
	_ "github.com/sirupsen/logrus"
	"go-server/middleware"
	"io"
	_ "net/http/pprof"
	"os"
)

const port = 20080
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
	//性能测试
	go func() {
		fmt.Println("pprof start")
		fmt.Println(http.ListenAndServe(":9876",nil))
	}()

	//start
	if err := router.Run(fmt.Sprintf(":%s",listen)); err != nil {
		println("Error when running server. " + err.Error())
	}

}