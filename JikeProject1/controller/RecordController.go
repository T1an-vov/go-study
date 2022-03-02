package controller

import (
	"JikeProject1/service"
	"github.com/gin-gonic/gin"
)

func RecordCreate(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	if err:=service.RecordCreate(email,number);err!=nil{
		c.JSON(200,gin.H{
			"message":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"message":"加入房间成功",
		})
	}
}

func RecordDelete(c *gin.Context)  {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	if err:=service.RecordDelete(email,number);err!=nil{
		c.JSON(200,gin.H{
			"message":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"message":"退出房间成功",
		})
	}
}
func RecordGet(c *gin.Context)  {
	number:=c.Query("number")
	records:=service.RecordGet(number)
	c.JSON(200,records)
}

func RecordDeleteRoom(c *gin.Context)  {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	if err:=service.RecordDeleteRoom(email,number);err!=nil{
		c.JSON(200,gin.H{
			"message":err.Error(),
		})
	}else {
		c.JSON(200,gin.H{
			"message":"删除成功",
		})
	}
}

func RecordDeleteUser(c *gin.Context)  {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	userEmail:=c.PostForm("userEmail")
	if err:=service.RecordDeleteUser(email,number,userEmail);err!=nil{
		c.JSON(200,gin.H{
			"message":err.Error(),
		})
	}else {
		c.JSON(200,gin.H{
			"message":"删除成功",
		})
	}
}