package service

import (
	"JikeProject1/dao"
	"JikeProject1/module"
	"errors"
)

func UserPost(user *module.User) error {
	if err:=dao.DB.Create(user).Error;err!=nil{
		return errors.New("邮箱已被使用，创建失败")
	}
	return nil
}
func UserDelete(email string) error {
	user:=new(module.User)
	if err:=dao.DB.Where("email=?",email).First(user).Error;err!=nil{
		return errors.New("该用户不存在")
	}
	if err:=dao.DB.Delete(user).Error;err!=nil{
		return err
	}
	return nil
}

func UserLoginIn(name string,password string) (*module.User,error) {
	user:=new(module.User)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(user).Error;err!=nil{
		return nil,errors.New("账号或密码错误")
	}
	return user,nil
}

func UserChangeName(email string, newName string) error {
	user:=new(module.User)
	if err:=dao.DB.Model(user).Where("email=?",email).Update("name",newName).Error;err!=nil{
		return err
	}
	return nil
}

func UserChangePassword(email string, newPassword string) error {
	user:=new(module.User)
	if err:=dao.DB.Model(user).Where("email=?",email).Update("password",newPassword).Error;err!=nil{
		return err
	}
	return nil
}

func UserChangeEmail(email string, newEmail string) error {
	user:=new(module.User)
	if err:=dao.DB.Model(user).Where("email=?",email).Update("email",newEmail).Error;err!=nil{
		return err
	}
	return nil
}

func  UserCheck(email string)bool{
	user:=new(module.User)
	if err:=dao.DB.Where("email=?",email).First(user).Error;err!=nil{
		return false
	}
	return true
}

func UserChangePasswordByEmail(email string,newPassword string){
	user:=new(module.User)
	dao.DB.Model(user).Where("email=?",email).Update("password",newPassword)
}


