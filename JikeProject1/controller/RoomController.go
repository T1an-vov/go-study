package controller

import (
	"JikeProject1/helper"
	"JikeProject1/module"
	"JikeProject1/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func RoomCreate(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	name:=c.PostForm("name")
	access:=c.PostForm("access")
	max, maxErr :=strconv.Atoi(c.PostForm("max"))
	if maxErr !=nil||max<1{
		c.JSON(200,gin.H{
			"message":"最大人数输入错误",
		})
		return
	}
	if access!="yes"&&access!="no"{
		c.JSON(200,gin.H{
			"message":"权限设置错误",
		})
		return
	}
	room:=&module.Room{
		Number:   helper.RandomDigitString(10),
		Name:     name,
		Owner:    email,
		Max:      uint(max),
		CreateAt: time.Now().Unix(),
		Access:   access,
		People:   1,
	}
	if CreateErr:=service.RoomCreate(room);CreateErr!=nil{
		c.JSON(200,gin.H{
			"message":"创建失败",
		})
	}else {
		c.JSON(200,gin.H{
			"message":"创建房间成功,房间号为："+room.Number,
		})
	}
}

func RoomDelete(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	flag,DeleteErr:=service.RoomDelete(number,email)
	if DeleteErr!=nil{
		c.JSON(200,gin.H{
			"message":"该房间不存在",
		})
	}else if !flag{
		c.JSON(200,gin.H{
			"message":"只有房主能删除房间",
		})
	}else {
		c.JSON(200,gin.H{
			"message":"删除成功",
		})
	}
}

func RoomChangeName(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	room:=service.RoomCheck(number)
	if room==nil{
		c.JSON(200,gin.H{
			"message":"该房间不存在",
		})
		return
	}
	if room.Owner!=email{
		c.JSON(200,gin.H{
			"message":"只有房主可以改名",
		})
		return
	}
	newName:=c.PostForm("newName")
	if err:=service.RoomChangeName(number,newName);err!=nil{
		c.JSON(200,gin.H{
			"message":"修改房间名失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"修改房间名成功",
	})
}

func RoomChangeMax(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	room:=service.RoomCheck(number)
	if room==nil{
		c.JSON(200,gin.H{
			"message":"该房间不存在",
		})
		return
	}
	if room.Owner!=email{
		c.JSON(200,gin.H{
			"message":"只有房主可以修改最大人数",
		})
		return
	}
	newMax,err:=strconv.Atoi(c.PostForm("newMax"))
	if err != nil {
		c.JSON(200,gin.H{
			"message":"最大人数输入错误",
		})
		return
	}
	if err=service.RoomChangeMax(number,newMax);err!=nil{
		c.JSON(200,gin.H{
			"message":"修改最大人数失败",
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"修改最大人数成功",
	})
}

func RoomChangeOwner(c *gin.Context) {
	email,CookieErr:=c.Cookie("user")
	if CookieErr!=nil{
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
		return
	}
	number:=c.PostForm("number")
	room:=service.RoomCheck(number)
	if room==nil{
		c.JSON(200,gin.H{
			"message":"该房间不存在",
		})
		return
	}
	if room.Owner!=email{
		c.JSON(200,gin.H{
			"message":"只有房主可以转让房主",
		})
		return
	}
	newOwner:=c.PostForm("newOwner")
	if !service.UserCheck(newOwner){
		c.JSON(200,gin.H{
			"message":"该用户不存在",
		})
		return
	}else {
		service.RoomChangeOwner(number, newOwner)
		c.JSON(200, gin.H{
			"message": "转让房主成功",
		})
	}
}

func RoomGet(c *gin.Context) {
	name:=c.Query("name")
	rooms:=service.RoomsGet(name)
	c.JSON(200,rooms)
}

