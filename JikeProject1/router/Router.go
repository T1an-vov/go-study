package router

import (
	"JikeProject1/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine{
	r:=gin.Default()
	UserGroup:=r.Group("user")
	{
		UserGroup.POST("/",controller.UserCreate)//创建用户
		Login:=UserGroup.Group("login")
		{
			Login.GET("welcome",controller.Welcome).DELETE("welcome",controller.Welcome)//初始欢迎
			Login.GET("/",controller.UserLoginIN)//登录
			Login.DELETE("/",controller.UserLoginOut)//退出登录
		}
		UserGroup.DELETE("/",controller.UserDelete)//删除用户
		Change:=UserGroup.Group("change")
		{
			Change.PUT("name",controller.UserChangeName)//修改名称
			Change.PUT("password",controller.UserChangePassword)//修改密码
			Change.PUT("email",controller.UserChangeEmail)//修改邮箱
			newChange:=Change.Group("p@ssword")
			{
				newChange.PUT("/",controller.UserChangePasswordByEmail1)
				newChange.PUT("ok",controller.UserChangePasswordByEmail2)
			}
		}//修改用户信息
		UserGroup.POST("img",controller.SendImg)
	}
	RoomGroup:=r.Group("room")
	{
		RoomGroup.POST("/",controller.RoomCreate)
		RoomGroup.DELETE("/",controller.RoomDelete)
		RoomGroup.GET("/",controller.RoomGet)
		RoomChangeGroup:=RoomGroup.Group("")
		{
			RoomChangeGroup.PUT("name",controller.RoomChangeName)
			RoomChangeGroup.PUT("max",controller.RoomChangeMax)
			RoomChangeGroup.PUT("owner",controller.RoomChangeOwner)
		}
	}
	RecordGroup:=r.Group("record")
	{
		RecordGroup.POST("/",controller.RecordCreate)
		RecordGroup.DELETE("/",controller.RecordDelete)
		RecordGroup.GET("/",controller.RecordGet)
		RecordRoomGroup:=RecordGroup.Group("room")
		{
			RecordRoomGroup.DELETE("all",controller.RecordDeleteRoom)
			RecordRoomGroup.DELETE("user",controller.RecordDeleteUser)
		}
	}
	return r
}