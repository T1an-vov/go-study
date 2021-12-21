package controller

import (
	"Work5/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecordsGet(c *gin.Context) {
	username := c.Query("username")
	if records, err := service.RecordsGet(username); err != nil||records==nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"records":records,
		})
	}
}
func RecordDelete(c *gin.Context) {
	username:=c.PostForm("username")
	course:=c.PostForm("coursename")
	if err:=service.RecordDelete(username,course);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":"删除"+username+"选课"+course,
		})
	}
}
func RecordPost(c *gin.Context) {
	username:=c.PostForm("username")
	coursename:=c.PostForm("coursename")
	if err:=service.RecordPost(username,coursename);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":400,
			"err":err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"message":username+"选课"+coursename+"成功",
		})
	}
}
