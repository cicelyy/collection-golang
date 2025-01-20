//创建user的服务
package service

import (
	"memo/model"
	"memo/pkg/utils"
	"memo/serializer"
	_"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

//封装用户注册和登录时需要的数据
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=16"`
}

//注册逻辑
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	//检查数据库中是否已存在相同用户名的用户
	model.DB.Model(&model.User{}).Where("user_name=?",service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response {
			Status : 400,
			Msg: "该用户已存在，无需再注册",
		}
	}
	user.UserName = service.UserName
	// 密码加密
	if err := user.SetPassword(service.Password);err != nil {
		return serializer.Response{
			Status: 400,
			Msg: err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg: "用户注册成功了",
	}
}

//处理用户登录逻辑 
func (service *UserService) Login() serializer.Response {
	var user model.User
	//先找user，看数据库中存不存在
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err){
			return serializer.Response{
				Status: 400,
				Msg: "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 400,
			Msg: "数据库错误",
		}
	}
	//对密码进行验证 
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg: "密码错误",
		}
	}
	//发一个token，为了其他功能需要身份所给前端存储的。
	//创建一个备忘录，这个功能就要token，不然不知道是谁创建的备忘录。
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg: "token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{User:serializer.BuildUser(user), Token:token},
		Msg : "登录成功",
	} 
} 