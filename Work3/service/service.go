package main

import (
	"Work3/proto"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
)

type GetUserReq struct{
	UserID int64 `json:"userId"`
}
type GetUserResp struct {
	UserID int64 `json:"userId"`
	UserName string `json:"userName"`
}
func process(conn net.Conn) {
	res:=new(GetUserReq)
	rep:=new(GetUserResp)
	defer conn.Close()
	reader:=bufio.NewReader(conn)
	num:=1//用于记录数据次数
	for{
		msg,err:=proto.Decode(reader)//将收到的信息解码
		if err == io.EOF {
			fmt.Println("读取完成")
			break
		}
		if err != nil {
			fmt.Println("decode failed, err = ", err)
			break
		}
		UnmarshalErr:=json.Unmarshal([]byte(msg),res)//json反序列化
		if UnmarshalErr != nil {
			fmt.Println("unmarshal failed")
			break
		}
		fmt.Println("收到第"+strconv.Itoa(num)+"次来自客户端的数据:")
		fmt.Println("userID:"+strconv.Itoa(int(res.UserID)))//打印收到的数据
		rep.UserID=res.UserID
		rep.UserName="第"+strconv.Itoa(num)+"个用户的名字"//将username设为"第x个用户的名字"
		num++
		repbyte,err:=json.Marshal(rep)//json序列化
		if err != nil {
			fmt.Println("rep marshal failed")
			break
		}
		b,err:=proto.Encode(string(repbyte))//编码
		if err != nil {
			fmt.Println("b encode failed")
			break
		}
		conn.Write(b)//发送信息
	}
}

func main() {
	lister,err:=net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed")
		return
	}
	defer lister.Close()
	for {
		conn,err:=lister.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}
		go process(conn)
	}
}
