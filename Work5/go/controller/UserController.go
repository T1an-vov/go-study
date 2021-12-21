package controller

import (
	"Work5/go/module"
	"Work5/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserPost(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	maxcredit,crediterr:=strconv.Atoi(c.PostForm("maxcredit"))
	if crediterr != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"err": "学分输入错误",
		})
		return
	}
	var user = module.User{
		Username:  username,
		Password:  password,
		Maxcredit: maxcredit,
	}
	err := service.UserPost(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"message":   "创建用户 " + user.Username,
		})
	}
}
func UserDelete(c *gin.Context) {
	username := c.PostForm("username")
	err := service.UseDelete(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"err": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除用户 " + username,
		})
	}
}
func UserGet(c *gin.Context) {
	username := c.Query("username")
	if user, err := service.UserGet(username); err != nil || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"message":   "查询 " + user.Username,
			"credit":    user.Credit,
			"maxcredit": user.Maxcredit,
		})
	}
}
func UserPut(c *gin.Context) {
	username := c.PostForm("username")
	newpassword := c.PostForm("newpassword")
	if err := service.UserPut(newpassword, username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "修改" + username + "的新密码为" + newpassword,
		})
	}
}
