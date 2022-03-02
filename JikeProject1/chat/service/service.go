package main

import (
	"fmt"
	"net"
)


type User struct {
	//名字
	name string
	//信息管道
	msg chan string
}


var Message = make(chan string, 10)//用于广播
var UsersMessage =make(map[string]User)//用于用户

func process(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("clientAddr:", clientAddr)
	//创建user
	newUser := User{
		name: clientAddr,
		msg:  make(chan string, 10),
	}
	//添加user到map结构
	UsersMessage[newUser.name] = newUser
	defer conn.Close() // 关闭连接
	go writeBackToClient(&newUser,conn)
	loginInfo := fmt.Sprintf("%s加入群聊\n", newUser.name,)
	Message <- loginInfo
	for {
		//具体业务逻辑
		buf := make([]byte, 1024)
		//读取客户端发送过来的请求数据
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("准备退出, err:", err)
			return
		}

		if err != nil {
			fmt.Println("conn.Read err:", err, ", cnt:", cnt)
			return
		}
		fmt.Println("服务器接收客户端发送过来的数据为: ", string(buf[:cnt]), ", cnt:", cnt)
		userInput := string(buf[:cnt])
		Message <- userInput
	}
}
func broadcast() {
	fmt.Println("广播启动成功...")
	defer fmt.Println("broadcast 程序退出!")
	for {
		m := <-Message
		for _, user := range UsersMessage {
			user.msg <- m
		}
	}
}

func writeBackToClient(user *User, conn net.Conn) {
	for data := range user.msg {
		fmt.Printf("user : %s 写回给客户端的数据为:%s\n", user.name, data)
		_, _ = conn.Write([]byte(data))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	go broadcast()//广播
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}