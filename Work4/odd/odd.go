package odd

import (
	"strconv"
)


//奇校验编码
func Parity(b byte) []byte {
	s:=strconv.FormatInt(int64(b),2)//先将字符转为字节然后将字节转为二进制数
	count:=0
	for i := 0; i < len(s); i++ {
		if s[i]=='1'{
			count++
		}
	}
	result:=[]byte(s)
	if count%2==1 {
		result=append(result,'0')
	}else {
		result=append(result,'1')
	}
	return result
}

//进行奇校验，如果校验成功同时解码
func Judge(b []byte) (bool,int64) {
	count:=0
	for _, b2 := range b {
		if b2 == '1' {
			count++
		}
	}
	if count%2==0 {
		return false,0
	}
	b=b[:len(b)-1]//删去最后一位奇校验码
	result,_:=strconv.ParseInt(string(b),2,64)
	return true,result
}
