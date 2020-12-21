package redisService

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

//分布式锁实现（不可重入）
type RedisLock struct {
	LockKey string
	value   string
}

//保证原子性
var delScript = `
if redis.call("get",KEY[1])==ARGV[1] then
	return redis.call("del",KEY[1])
else
	return 0
end
`

func (rl *RedisLock) Lock(client *redis.Client, timeout int) error {
	//随机数
	if client == nil {
		return errors.New("redis client is nil")
	}
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	rl.value = base64.StdEncoding.EncodeToString(b)
	ok, err := client.SetNX(rl.LockKey, rl.value, time.Duration(timeout)*time.Second).Result()
	if err == nil && ok {
		errMsg := fmt.Sprintf("lock filed")
		return errors.New(errMsg)
	} else if err == nil {
		return nil
	} else {
		errMsg := fmt.Sprintf("redis lock fail err: %v", err)
		return errors.New(errMsg)
	}
}

//解锁
func (rl *RedisLock) Unlock(client *redis.Client) error {
	if client == nil {
		return errors.New("redis client is nil")
	}
	_, err := client.Eval(delScript, []string{rl.LockKey, rl.value}).Result()
	return err
}