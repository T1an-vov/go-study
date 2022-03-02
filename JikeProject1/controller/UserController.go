package controller

import (
	"JikeProject1/helper"
	"JikeProject1/module"
	"JikeProject1/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var str string//验证码

//欢迎界面
func Welcome(c *gin.Context) {
	c.String(200,"欢迎")
}

//创建用户
func UserCreate(c *gin.Context) {
	var user module.User
	c.ShouldBind(&user)
	err := service.UserPost(&user)
	if user.Name==""||user.Password==""||user.Email==""{
		c.JSON(200,gin.H{
			"code":200,
			"message":"信息不能为空",
		})
		c.Abort()
	} else if err != nil{
		c.JSON(200, gin.H{
			"code":    200,
			"err": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":      200,
			"message":   "创建用户 " + user.Name,
		})
	}
}

//根据邮箱删除用户
func UserDelete(c *gin.Context) {
	email := c.PostForm("email")
	err := service.UserDelete(email)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    200,
			"err": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除用户 ",
		})
	}
}

//登录
func UserLoginIN(c *gin.Context) {
	name := c.Query("name")
	password:=c.Query("password")
	if user, err := service.UserLoginIn(name,password); err != nil {
		c.JSON(200, gin.H{
			"code":    200,
			"err": err.Error(),
		})
	} else {
		cookie:=&http.Cookie{
			Name:       "user",
			Value:      user.Email,
			MaxAge:     3600,
			Path:     "/",
			HttpOnly:   true,
		}
		http.SetCookie(c.Writer,cookie)
		//c.SetCookie("user",user.Email,3600,"/","localhost",false,true)
		c.JSON(200, gin.H{
			"code": 200,
			"message":user.Name+"登录成功",
		})
	}

}

//退出登录
func UserLoginOut(c *gin.Context) {
	//cookieValue,err:=c.Request.Cookie("user")
	cookieValue,err:=c.Cookie("user")
	if err != nil {
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
	}
	cookie:=&http.Cookie{
		Name:       "user",
		Value:      cookieValue,
		MaxAge:     -1,
		Path:     "/",
		HttpOnly:   true,
	}
	http.SetCookie(c.Writer,cookie)
	c.Redirect(http.StatusTemporaryRedirect,"welcome")
}

//修改名称
func UserChangeName(c *gin.Context) {
	email,err:=c.Cookie("user")
	//cookie,err:=c.Cookie("user")
	if err != nil {
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
	}
	newName := c.PostForm("newName")
	if err=service.UserChangeName(email,newName);err!=nil{
		c.JSON(200,gin.H{
			"message":"修改名称失败",
		})
	}else{
		c.JSON(200,gin.H{
			"message":"修改名称成功",
		})
	}
}


//修改密码
func UserChangePassword(c *gin.Context) {
	email,err:=c.Cookie("user")
	if err != nil {
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
	}
	newPassword := c.PostForm("newPassword")
	if err=service.UserChangePassword(email,newPassword);err!=nil{
		c.JSON(200,gin.H{
			"message":"修改密码失败",
		})
	}else{
		c.JSON(200,gin.H{
			"message":"修改密码成功",
		})
	}
}

//修改邮箱
func UserChangeEmail(c *gin.Context) {
	email,err:=c.Cookie("user")
	//cookie,err:=c.Cookie("user")
	if err != nil {
		c.JSON(200,gin.H{
			"message":"请先登录",
		})
	}
	newEmail := c.PostForm("newEmail")
	if err=service.UserChangeEmail(email,newEmail);err!=nil{
		c.JSON(200,gin.H{
			"message":"修改邮箱失败",
		})
	}else{//修改邮箱后cookie也会跟着变化
		cookie:=&http.Cookie{
			Name:       "user",
			Value:      newEmail,
			MaxAge:     3600,
			Path:     "/",
			HttpOnly:   true,
		}
		http.SetCookie(c.Writer,cookie)
		c.JSON(200,gin.H{
			"message":"修改邮箱成功",
		})
	}
}

//通过邮箱验证码修改密码:通过邮箱发送验证码
func UserChangePasswordByEmail1(c *gin.Context){
	email:=c.PostForm("email")//获取邮箱
	if !service.UserCheck(email){
		c.JSON(200,gin.H{
			"message":"该邮箱对应用户不存在",
		})
	}else {
		str= helper.GetRandomString() //生成验证码
		helper.SendEmail(email,str)   //发送验证码
	}
}
//通过邮箱验证码修改密码:比对验证码，正确则修改密码
func UserChangePasswordByEmail2(c *gin.Context)  {
	email:=c.PostForm("email")
	check:=c.PostForm("str")//从用户处接收验证码
	newPassword:=c.PostForm("newPassword")
	if str==check&&str!=""{//比对验证码
		service.UserChangePasswordByEmail(email,newPassword)
		str=""
		c.JSON(200,gin.H{
			"message":"修改密码成功",
		})
	}else{
		c.JSON(200,gin.H{
			"message":"验证码错误",
		})
	}
	//
}
//
func SendImg(c *gin.Context)  {
	file,err:=c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"err:":err.Error(),
		})
		return
	}
	dst:=fmt.Sprintf("./%s",file.Filename)
	c.SaveUploadedFile(file,dst)
}
