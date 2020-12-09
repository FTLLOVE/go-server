package redisService

import (
	"errors"
	"fmt"
	"go-server/data"
	"log"
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
	userHasCouponKey := getHasCouponKeyByUser(userName)
	couponKey:=getCouponKeyByName(couponName)
	res,err:=data.EvalSHA(secKillSHA,[]string{userHasCouponKey,couponName,couponKey})
	if err!=nil{
		return -1,errors.New("eval error")
	}
	couponRes,ok :=res.(int64)
	if !ok{
		return -1, errors.New("type error")
	}
	//不同的错误提示信息
	switch {
	case couponRes==-1:
		return -1,errors.New("hasCoupon")
	case couponRes==-2:
		return -1,errors.New("noCoupon")
	case couponRes==-3:
		return -1,errors.New("noRemain")
	case couponRes==1:
		return couponRes,nil
	default:
		{
			log.Fatal("Unexpected return value")
			return -1,errors.New("no return value")
		}
	}
}

