package gin_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/Service/redisService"
	"go-server/basic"
	"go-server/data"
	"go-server/model"
	"log"
	"net/http"
)

const (
	MinUserNameLen = 1
	MinPasswordLen = 1
	NormalCustomer = 1
	NormalSeller   = 2
)

type FetchCouponReq struct {
	UserName string `json:"user_name" form:"user_name"`
	CouponName string `json:"coupon_name" form:"coupon_name"`
}

type FetchCouponRsp struct {
	Rsp struct{
		Description string `json:"description"`
		Status 		string `json:"status"`
		Data struct{
			
		} `json:"data"`
	}
}

//登陆middleware处理
//失败返回错误信息，成功通过channel异步落库
func FetchCoupon(c*gin.Context)  {
	req:=FetchCouponReq{}
	if err:=c.Bind(&req);err!=nil{
		fmt.Println("FetchCoupon prarm err",err.Error())
	}
	_,err:=redisService.AtomicSecKill(req.UserName,"sellname",req.CouponName)
	if err!=nil{
		coupon:=redisService.GetCoupon(req.CouponName)
		secKillChannel<-secKilMessage{req.UserName,coupon}
		c.JSON(http.StatusCreated, gin.H{"err": ""})
		return
	}else{
		if redisService.IsEvalError(err) {
			log.Printf("Server error" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"evalErr": err.Error()})
			return
		} else {
			//log.Println("Fail to fetch coupon. " + err.Error())
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
}

type RegisterUserReq struct {
	UserName string `json:"user_name" form:"UserName" binding:"required"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required"`
	Kind  	 int `json:"kind" form:"kind"` //买家or卖家
}

type RegisterUserRsp struct {
	Status string `json:"status"`
	Description string `json:"description"`
	Data struct{
		
	} `json:"data"`
}

//用户注册
func RegisterUser(c *gin.Context) {
	req := RegisterUserReq{}
	rsp := RegisterUserRsp{}.Data
	if err := c.Bind(&req); err != nil {
		log.Println("RegisterUser param error", err.Error())
		basic.ResponseError(c, rsp, err.Error())
		return
	}
	if len(req.UserName) < MinUserNameLen {
		log.Println("user name length too short")
		basic.ResponseError(c, rsp, "name too short")
		return
	}
	if len(req.PassWord) < MinPasswordLen {
		log.Println("password length too short")
		basic.ResponseError(c, rsp, "password")
		return
	}
	if req.Kind == 0 {
		log.Println("kind is error")
		basic.ResponseError(c, rsp, "kind is error")
		return
	}

	//加入用户
	req.PassWord = model.GetMd5(req.PassWord) //MD5加密，base64编码方便网络传输
	err := new(model.User).Insert(data.Db, req.UserName, req.PassWord, req.Kind)
	if err != nil {
		log.Println("create User error", err.Error())
		basic.ResponseError(c, rsp, "user Register error"+err.Error())
		return
	}
	basic.ResponseOk(c, rsp, "OK")
	return
}