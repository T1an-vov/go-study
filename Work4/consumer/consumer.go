package consumer

import (
	"Work4/odd"
	"Work4/proto"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func Consumer(ch chan string)  {
	result:=[]rune{}//用于保存最终的结果
	for i:=range ch {
		time.Sleep(1 * time.Second)
		result=result[0:0]//清空
		fmt.Println("收到发送的数据："+i)
		file, err := os.Open("./mq/mq1.mq")
		if err != nil && err != io.EOF {
			fmt.Println("文件打开错误"+err.Error())
		}
		defer file.Close()
		read := bufio.NewReader(file)
		for {
			msg, err := proto.Decode(read)//粘包解码
			flag,b:= odd.Judge([]byte(msg))//奇校验解码
			if !flag&&len(msg)>0 {
				fmt.Printf("奇校验失败,err:%v",err)
				break
			}
			if err==io.EOF{
				break
			}
			if err != nil {
				fmt.Printf("解码失败:%v\n", err)
				break
			}
			result=append(result,rune(b))
		}
	}
	fmt.Printf("读取到文件中内容为：%s\n", string(result))
}