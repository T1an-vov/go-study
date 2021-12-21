package controller

import (
	"Work5/go/module"
	"Work5/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CoursePost(c *gin.Context)  {
	var course module.Course
	c.ShouldBind(&course)
	err:=service.CoursePost(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"创建课程 "+course.Coursename,
		})
	}
}
func CourseGet(c *gin.Context) {
	var course  *module.Course
	var err error
	coursename:=c.Query("coursename")
	if course,err=service.CourseGet(coursename);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":course.Coursename,
			"number":course.Number,
			"maxnumber":course.Maxnumber,
			"credit":course.Credit,
		})
	}
}
func CoursePut(c *gin.Context) {
	var course module.Course
	c.ShouldBind(&course)
	if err:=service.CoursePut(course.Coursename,&course);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"更新课程"+course.Coursename+"成功",
		})
	}
}
func CourseDelete(c *gin.Context) {
	coursename:=c.PostForm("coursename")
	if err:=service.CourseDelete(coursename);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"删除"+coursename+"成功",
		})
	}
}