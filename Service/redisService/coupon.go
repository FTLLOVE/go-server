package redisService

import (
	"fmt"
	"go-server/data"
	"go-server/model"
	"strconv"
)

//暂时使用name 实际使用id更合适 使用hash存储优惠券信息待调整

//获取已有的coupon
func getHasCouponKeyByUser(name string) string {
	return fmt.Sprintf("%s-has", name)
}

//获取优惠券key用于缓存预热
func getCouponKeyByName(name string) string {
	return fmt.Sprintf("%s-info", name)
}

//缓存被用户拥有的优惠券
func CacheHasCoupon(coupon model.Coupon) (int64, error) {
	key := getHasCouponKeyByUser(coupon.UserName)
	val, err := data.SetAdd(key, coupon.CouponName)
	return val, err
}

//缓存优惠券完整信息
func CacheCoupon(coupon model.Coupon) (bool, error) {
	key := getCouponKeyByName(coupon.CouponName)
	fileds := map[string]interface{}{
		"id":          coupon.Id,
		"username":    coupon.UserName,
		"couponName":  coupon.CouponName,
		"amount":      coupon.Amount,
		"remain":      coupon.Remain,
		"stock":       coupon.Stock,
		"description": coupon.Description,
	}
	val, err := data.SetHash(key, fileds)
	return val, err
}

//缓存获取优惠券信息
func GetCoupon(couponName string) model.Coupon {
	key := getCouponKeyByName(couponName)
	values, err := data.GetMap(key, "id", "username", "couponName", "amount", "remain", "stock", "description")
	if err != nil {
		fmt.Println("Error on get coupon" + err.Error())
	}
	id, err := strconv.ParseInt(values[0].(string), 10, 64)
	if err != nil {
		fmt.Println("Wrong type of id" + err.Error())
	}
	amount, _ := strconv.ParseInt(values[3].(string), 10, 64)

	remain, _ := strconv.ParseInt(values[4].(string), 10, 64)
	stock, _ := strconv.ParseInt(values[5].(string), 10, 64)
	return model.Coupon{
		Id:          id,
		UserName:    values[1].(string),
		CouponName:  values[2].(string),
		Amount:      amount,
		Remain:      remain,
		Stock:       stock,
		Description: values[6].(string),
	}
}

//缓存获取用户所有优惠券
func GetCouponsByUser(userName string) ([]model.Coupon, error) {
	var coupons []model.Coupon
	hasKey := getHasCouponKeyByUser(userName)
	couponName, err := data.GetMembers(hasKey)
	if err != nil {
		fmt.Println("error in getmembers coupon" + err.Error())
		return coupons, err
	}
	for _, v := range couponName {
		cou := GetCoupon(v)
		coupons = append(coupons, cou)
	}
	return coupons, nil
}
