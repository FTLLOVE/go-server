package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

type HelloReq struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type HelloRsp struct {
	Rsp struct{
		Status string `json:"status"`
		Description string `json:"description"`
		Data struct{
			Name string `json:"name"`
		} `json:"data"`
	}
	
}
func GetTest(c*gin.Context)  {
	req:= HelloReq{}
	rsp:= HelloRsp{}
	if err:=c.Bind(&req);err!=nil{
		log.Fatal("tset")
	}
	rsp.Rsp.Data.Name = req.Name
	c.JSON(200,rsp)
}
