package gin_server

import (
	"go-server/model"
	"log"
)

type secKilMessage struct {
	username string
	coupon  model.Coupon
}

const  maxMessageNum = 20000

var secKillChannel = make(chan secKilMessage,maxMessageNum)

//异步存储
func Consumer()  {

	for{
		message := <-secKillChannel
		log.Println("get message: ",message)
		//用户优惠添加

		//优惠券总库存减少

	}

}

var cousumerRun = false

func RunConsumer()  {
	if !cousumerRun{
		go Consumer()
		cousumerRun = true
	}
}
