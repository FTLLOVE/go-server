package Redis

import (
	"github.com/go-redis/redis/v7"
	"go.mod/configInit"
	"log"
)
var client *redis.Client

func redisInit(config configInit.Config)  {
	client = redis.NewClient(&redis.Options{
		Addr: config.Redis.Address,
		Password: config.Redis.Password,
		DB: 0,
	})

	_,err:=FlushAll()
	if err!=nil{
		log.Fatal("flushAll",err.Error())
	}

}

func FlushAll()(string,error)  {
	return client.FlushAll().Result()
}

//加载lua脚本
func LoadScript(script string)string  {
	scriptExists,err:= client.ScriptExists(script).Result()
	if err!=nil{
		panic("exist failed err: %v"+err.Error())
	}
	if !scriptExists[0]{
		scriptSHA,err:=client.ScriptLoad(script).Result()
		if err!=nil{
			panic("load script error"+err.Error())
		}
		return scriptSHA
	}
	log.Println("script exists")
	return ""
}

func EvalSHA(SHA string,args[]string)(interface{},error)  {
	val,err:=client.Eval(SHA,args).Result()
	if err!=nil{
		log.Println("eval failed err",err.Error())
		return nil, err
	}
	return val,err
}

