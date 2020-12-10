package gin_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/Service/redisService"
	"go-server/model"
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
	req:=FetchCouponRsp{}
	if err:=c.Bind(&req);err!=nil{
		fmt.Println("FetchCoupon prarm err",err.Error())
	}
	_,err:=redisService.AtomicSecKill(req.UserName,"sellname",req.CouponName)
	if err!=nil{
		coupon:=redisService.GetCoupon(req.CouponName)

	}
}