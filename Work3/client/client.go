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
func main() {
	var user =GetUserReq{}
	conn,err:=net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("client failed")
		return
	}
	defer conn.Close()
	rep:=new(GetUserResp)
	for i:=1;i<=20;i++{//一共进行20次数据收发操作，i用于记录次数
		user.UserID=int64(i)//将userID设为i
		JsonUser,err:=json.Marshal(user)//json序列化
		if err != nil {
			fmt.Println("client json err")
		}
		b,err:=proto.Encode(string(JsonUser))//编码
		if err != nil {
			fmt.Println("client encode err")
		}
		conn.Write(b)//发送数据
		reader:=bufio.NewReader(conn)
		msg,err:=proto.Decode(reader)//将收到的数据解码
		if err == io.EOF {
			fmt.Println("读取完成")
			break
		}
		if err != nil {
			fmt.Println("decode failed")
			return
		}
		UnmarshalErr:=json.Unmarshal([]byte(msg),rep)//反序列化
		if UnmarshalErr != nil {
			fmt.Println("unmarshal failed")
			continue
		}
		fmt.Println("收到第"+strconv.Itoa(i)+"次来自服务端的数据:")
		fmt.Println("userID:"+strconv.Itoa(int(rep.UserID)))
		fmt.Println("userName:"+rep.UserName)//打印收到的数据
	}
}