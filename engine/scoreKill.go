package engine

import (
	"fmt"
)

//定义返回error

type evalError struct {

}

func (e evalError)Error()string  {
	return "executing redis eval error"
}

type hasCouponError struct {
	userName string
	couponName string
}

//优惠券已经拥有
func (e hasCouponError)Error()string  {
	return fmt.Sprintf("User %s has coupon %s",e.userName,e.couponName)
}


type noneCouponError struct {
	userName string
	couponName string
}

//用户优惠券不存在
func (e noneCouponError)Error()string  {
	return fmt.Sprintf("%s coupon %s is not exist",e.couponName,e.userName)
}

func IsEvalError(err error)bool{
	switch err.(type) {
	case evalError:
		return true
	default:
		return false
	}
}


//执行redis原子秒杀
func AtomicSecKill(userName string,sellName string,couponName string)(int64,error)  {
	//预先加载原子性的lua脚本

}

