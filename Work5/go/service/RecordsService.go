package service

import (
	"Work5/go/dao"
	"Work5/go/module"
	"errors"
)

func RecordsGet(username string) ([]module.Record, error){
	var records []module.Record
	if err:=dao.DB.Where("username=?",username).Find(&records).Error;err!=nil||len(records)==0{
		return nil,errors.New("记录不存在，查询失败")
	}
	return records,nil
}
func RecordPut()  {
	//选课记录只有查询、删除和创建
}

func RecordPost(username string,coursename string)(err error)  {
	var record module.Record
	user:=new(module.User)
	course:=new(module.Course)
	if err=dao.DB.Where("username=?",username).First(user).Error;err!=nil{
		return errors.New("用户不存在")
	}//查询是否有这个人
	if err=dao.DB.Where("coursename=?",coursename).First(course).Error;err!=nil{
		return errors.New("课程不存在")
	}//查询是否有这个课
	if err=dao.DB.Where(&module.Record{Username:username, Coursename:coursename}).First(&record).Error;err==nil{
		return errors.New(username+"已选过"+coursename)
	}
	if (user.Maxcredit-user.Credit)<course.Credit||course.Maxnumber==course.Number{
		return  errors.New("学分超过上限或者选课人数已满，选课失败")
	}//若学分超过上限或者选课人数已满，则选课失败
	user.Credit+=course.Credit
	course.Number+=1
	if err=dao.DB.Save(user).Error;err!=nil{
		return err
	}
	if err=dao.DB.Save(course).Error;err!=nil{
		return err
	}//保存选课后的结果
	if err=dao.DB.Create(&module.Record{Username:username,Coursename:coursename}).Error;err!=nil{
		return err
	}
	return nil
}

func RecordDelete(username string,coursename string) (err error) {
	var record module.Record
	course:=new(module.Course)
	user:=new(module.User)
	if err=dao.DB.Where(&module.Record{Username:username, Coursename:coursename}).First(&record).Error;err!=nil{
		return errors.New("该选课记录不存在，删除失败")
	}//先查询是否有这个选课记录
	if err=dao.DB.Where("username=?",username).First(user).Error;err!=nil{
		return err
	}//取得用户
	if err=dao.DB.Where("coursename=?",coursename).First(course).Error;err!=nil{
		return err
	}//取得课程
	user.Credit=user.Credit-course.Credit//改变用户学分
	course.Number-=1//改变课程选课人数
	if err=dao.DB.Save(user).Error;err!=nil{
		return err
	}
	if err=dao.DB.Save(course).Error;err!=nil{
		return err
	}//保存删除选课后的结果
	if err=dao.DB.Delete(&record).Error;err!=nil{
		return err
	}//删除记录
	return nil
}
