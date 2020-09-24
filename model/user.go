package model

import (
	"crypto/md5"
	"encoding/hex"
	_"github.com/jinzhu/gorm"
	_"os"
)

const (
	NormalUser = 1
	SellerUser = 2
)

type LoginUser struct {
	Name string
	Password string
}

//用户和商家注册类型不同
type RegisterUser struct {
	LoginUser
	Kind int
}

type User struct{
	Id int `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Kind int `json:"kind" gorm:"kind"`
}

func (user User)IsNormalUser()bool  {
	return user.Kind  == NormalUser
}

func (user User)IsSeller()bool  {
	return user.Kind == SellerUser
}

//用户类型格式
func IsValidUser(kind int)bool  {
	if kind==NormalUser||kind==SellerUser{
		return true
	}else{
		return false
	}
}

//md5 加密
func GetMd5(text string)string  {
	//md5其它方式
	//方法一
	/*hash:=md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))*/

	hash:=md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}