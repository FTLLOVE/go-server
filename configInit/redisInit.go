package configInit

import (
	"github.com/go-redis/redis/v7"
	"log"
)
var client *redis.Client

func redisInit(config Config)  {
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

//set time forever
func SetTime(key string,values interface{})(string,error)  {
	val,err:=client.Set(key,values,0).Result()
	return val,err
}

//设置hash值
func SetHash(key string,field map[string]interface{})(bool,error) {
	return client.HMSet(key,field).Result()
}

//获取hash表多个字段值
func GetMap(key string,fields ...string) ([]interface{},error) {
	return client.HMGet(key,fields...).Result()
}


func SetAdd(key string,field string)(int64,error)  {
	return  client.SAdd(key,field).Result()
}

//判断是否是集合的值
func SetIsMember(key string,field string)(bool,error)  {
	return client.SIsMember(key,field).Result()
}

//获取集合所有成员
func GetMembers(key string)([]string,error)  {
	return client.SMembers(key).Result()
}
