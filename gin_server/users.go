package gin_server

import "github.com/gin-gonic/gin"

type FetchCouponReq struct {
	UserName string `json:"user_name" form:"user_name"`
	CouponName string `json:"coupon_name" form:"coupon_name"`
}
func FetchCoupon(c*gin.Context)  {
	req:=FetchCouponReq{}
}