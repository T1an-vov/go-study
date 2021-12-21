package service

import (
	"Work5/go/dao"
	"Work5/go/module"
	"errors"
)

func CoursePost(course *module.Course) (err error) {
	if err=dao.DB.Create(&course).Error;err!=nil{
		return errors.New("课程已经存在，创建失败")
	}
	return nil
}
func CourseGet(coursename string)(*module.Course , error)  {
	var course module.Course
	if err:=dao.DB.Where("coursename=?",coursename).First(&course).Error;err!=nil{
		return nil,errors.New("课程不存在，查询失败")
	}
	return &course,nil
}
func CoursePut(coursename string,newcourse *module.Course) error {
	var course module.Course
	if err:=dao.DB.Where("coursename=?",coursename).First(&course).Error;err!=nil{
		return errors.New("课程不存在，更新失败")
	}
	if err:=dao.DB.Model(&course).Update(newcourse).Error;err!=nil{
		return err
	}
	return nil
}

func CourseDelete(coursename string)(err error){
	var course module.Course
	if err=dao.DB.Where("coursename=?",coursename).First(&course).Error;err!=nil{
		return errors.New("课程不存在,删除失败")
	}
	if err=dao.DB.Delete(&course).Error;err!=nil{
		return err
	}
	return nil
}