package service

import (
	"Work5/go/dao"
	"Work5/go/module"
	"errors"
)

func UserPost(user *module.User) (err error) {
	if user.Maxcredit==0{
		return errors.New("请输入正确的最大学分")
	}
	if user.Password==""{
		return errors.New("密码不能为空")
	}
	if err = dao.DB.Create(user).Error; err != nil {
		return errors.New("该用户已存在，创建失败")
	}
	return nil
}//传入user结构体保存进user表中
func UseDelete(username string) (err error) {
	var user module.User
	if err = dao.DB.Where("username=?", username).First(&user).Error; err != nil {
		return errors.New("用户不存在，删除失败")
	}//先查询是否已有此人
	if err = dao.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
func UserGet(username string) ( *module.User, error) {
	var user module.User
	if err := dao.DB.Where("username=?", username).First(&user).Error; err != nil {
		return nil, errors.New("用户不存在，查询失败")
	}
	return &user, nil
}//返回user结构体指针
func UserPut(newpassword string, username string) (err error) {
	var user module.User
	if err = dao.DB.Where("username=?", username).First(&user).Error; err != nil {
		return errors.New("用户不存在，修改密码失败")
	}//先查询是否存在此人
	user.Password = newpassword
	if err = dao.DB.Save(&user).Error; err != nil {
		return err
	}//修改密码
	return nil
}
