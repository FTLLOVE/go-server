package gin_server

import (
	"github.com/gin-gonic/gin"
)
func StartHttpServer(listen string)  {

	router:=gin.Default()//暂时使用default
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