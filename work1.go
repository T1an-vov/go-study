package main
import "fmt"
func main() {
	var ch1=make(chan string)
	var ch2 = make(chan string)
	var ch3 = make(chan string)
	var ch4 = make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			ch1<-"张三"
		}()
		go func() {
			ch2<-"李四"
		}()
		go func() {
			ch3<-"王五"
		}()
		go func() {
			ch4<-"赵六"
		}()
	}
	for i:=0;i<10;i++{
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
		fmt.Println(<-ch3)
		fmt.Println(<-ch4)
	}
}

