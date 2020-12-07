package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

//redis配置
var configFile []byte

const path string = "config-dev.yaml"

type Config struct {
	Redis Redis `json:"redis"`
}

type Redis struct {
	Address string `json:"address"`
	Network string `json:"network"`
	Password string `json:"password"`
	MaxIdle string `json:"maxidle"`
	MaxActive string `json:"maxactive"`
	Timeout string `json:"timeout"`
}

//获取配置信息
func GetConfig()(config Config,err error)  {
	err = yaml.Unmarshal(configFile,&config)
	return
}

//初始时自动调用加载redis信息到比特流
func init()  {
	var err error
	configFile,err = ioutil.ReadFile(path)
	if err!=nil{
		log.Fatal("redis init file open err : %v",err)
	}

}