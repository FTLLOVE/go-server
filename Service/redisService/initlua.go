package redisService

import (
	"go-server/data"
)

const secKillScript = `
	--keys 1 2 3 外部传递用于脚本使用
	--keys[1] :hasCouponKey "{username}-has"
	--keys[2]: couponName   "{couponName}"
    --keys[3]: couponKey    "{couponName}-info"
	--返回-1 -2 -3 失败返回不同错误 1 为成功
	
	--检查优惠券是否存在
	local couponLeft = redis.call("hget",KEYS[3],"left");
	if(couponLeft==false)
	then
		return -2; --不存在
	end

	if(tonumber(couponLeft)==0)
	then
		return -3 --优惠券无库存
	end

	local userHasCoupon = redis.call("SISMEMBER",KEYS[1],KEYS[2]);
	if(userHasCoupon ==1 ) --用户已经有此优惠券
	then
		return -1;
	end

	--到此处没错误 用户得到优惠券 返回1
	redis.call("hset",KEYS[3],"left",couponLeft-1);
	redis.call("SADD",KEYS[1],KEYS[2]);
	return 1;
`
//保存lua脚本的sha值
var secKillSHA string

//优惠券对应数据加载到缓存，防止缓存穿透
func preHeatKeys()  {

}

func init()  {
	//加载lua脚本
	secKillSHA  = data.LoadScript(secKillScript)

	//预热
	preHeatKeys()
}