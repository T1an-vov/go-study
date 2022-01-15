package main

import (
	"Work4/consumer"
	"Work4/producer"
	"fmt"
	"os"
)
func main() {
	ch:=make(chan string)
	go producer.Producer(ch)
	consumer.Consumer(ch)
	RemoveErr := os.Remove("./mq/mq1.mq")
	if RemoveErr != nil {
		fmt.Printf("删除文件失败:%v\n", RemoveErr)
	}else {
		fmt.Println("删除文件")
	}
}
