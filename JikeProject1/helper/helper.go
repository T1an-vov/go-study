package helper

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)


//生成随机验证码用于用户找回密码
func  GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//发送邮件
func SendEmail(email string,str string)error{
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "981774028@qq.com")
	//接收人
	m.SetHeader("To", email)
	//主题
	m.SetHeader("Subject", "修改密码验证码")
	//内容
	m.SetBody("text/html", "修改密码的验证码为:"+str)
	//附件
	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "981774028@qq.com", "yljwokpyjkkvbbcd")
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}



//产生随机数字字符串用于设置房间号
func RandomDigitString(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(48, 57))
	}
	return string(bytes)
}
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}