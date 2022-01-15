package producer

import (
	"Work4/odd"
	"Work4/proto"
	"fmt"
	"os"
)

func Producer(ch chan<- string) {
	var str string
	for n:=1;n<=3;n++{//总共进行三次收发
		fmt.Scan(&str)
		ch <- str
	for i:= 0; i < len(str); i++ {
		b := odd.Parity(str[i])
		msg, err := proto.Encode(string(b))
		if err != nil {
			fmt.Printf("encode failed,err:%v", err)
		}
		file, err := os.OpenFile("./mq/mq1.mq", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("open file failed, err:", err)
			return
		}
		defer file.Close()
		file.Write(msg)//写入内容
	}
	}
	close(ch)
}