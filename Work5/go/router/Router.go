package router

import (
	"Work5/go/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	UserGroup := r.Group("user")
	{
		UserGroup.POST("/", controller.UserPost)
		UserGroup.DELETE("/", controller.UserDelete)
		UserGroup.GET("/", controller.UserGet)
		UserGroup.PUT("/",controller.UserPut)
	}
	CourseGroup:=r.Group("course")
	{
		CourseGroup.POST("/",controller.CoursePost)
		CourseGroup.GET("/",controller.CourseGet)
		CourseGroup.PUT("/",controller.CoursePut)
		CourseGroup.DELETE("/",controller.CourseDelete)
	}
	RecordGroup:=r.Group("record")
	{
		RecordGroup.GET("/",controller.RecordsGet)
		RecordGroup.DELETE("/",controller.RecordDelete)
		RecordGroup.POST("/",controller.RecordPost)
	}
	return r
}
