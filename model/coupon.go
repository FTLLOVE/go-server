package model

type Coupon struct {
	Id         int64  `json:"id" gorm:"id"`
	UserName   string `json:"user_name" gorm:"user_name"`
	CouponName string `json:"coupon_name" gorm:"coupon_name"`
	//最大优惠券数量
	Amount int64 `json:"amount" gorm:"amount"`
	//剩余数量
	Remain int64 `json:"remain" gorm:"remain"`
	//优惠券面额
	Stock int64 `json:"stock" gorm:"stock"`
	//优惠券信息
	Description string `json:"description" gorm:"description"`
}

type CouponCommon struct {
	Name        string
	Stock       int64
	Description string
}

//商家查询数据
type SellCoupon struct {
	CouponCommon
	Amount int64
	Remain int64
}

//消费者查询返回数据
type CustomerCoupon struct {
	CouponCommon
}

//返回商家结构
func DtoSellCoupon(coupon []Coupon) []SellCoupon {
	var sellCoupon []SellCoupon
	for _, cou := range coupon {
		sellCoupon = append(sellCoupon, SellCoupon{
			CouponCommon: CouponCommon{
				Name:        cou.UserName,
				Stock:       cou.Stock,
				Description: cou.Description,
			},
			Amount: cou.Amount,
			Remain: cou.Remain,
		})
	}
	return sellCoupon
}

//返回消费者结构
func DtoCustomerCoupon(coupon []Coupon) []CustomerCoupon {
	var custCoupon []CustomerCoupon
	for _, cou := range coupon {
		custCoupon = append(custCoupon, CustomerCoupon{
			CouponCommon{
				cou.UserName,
				cou.Stock,
				cou.Description,
			},
		})
	}

	return custCoupon
}