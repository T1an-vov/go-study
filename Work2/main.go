package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file,err:=os.Open("top.txt")
	if err != nil {
		fmt.Println("open failed")
		return
	}//打开top文件
	defer file.Close()
	CpuData,err:=os.OpenFile("CpuData.txt",os.O_RDWR,0666)
	if err != nil {
		fmt.Println("open cpu data failed")
		return
	}//打开cpudata文件并准备写入
	defer CpuData.Close()
	MemData,err:=os.OpenFile("MemData.txt",os.O_RDWR,0666)
	if err != nil {
		fmt.Println("open mem data failed")
		return
	}//打开memdata文件并准备写入
	defer MemData.Close()
	reader:=bufio.NewReader(file)
	CpuRe1:=regexp.MustCompile("\\d+\\.\\d+ id")//cpu正则匹配1，用于匹配xx.xx id
	CpuRe2:=regexp.MustCompile("\\d+\\.\\d+")//cpu正则匹配2，用于匹配出xx.xx
	MemRe1:=regexp.MustCompile("MiB Mem.+used")//mem正则匹配1，用于匹配出MiB....used行
	MemRe2:=regexp.MustCompile("\\d+\\.\\d+ used")//mem正则匹配2，用于匹配xxx.xx used
	MemRe3:=regexp.MustCompile("\\d+\\.\\d+")//mem正则匹配3，用于匹配xx.xx
	var (
		i = 1
		j=1
	)//两个计数器，记录读取了多少组数据

	for{
		var (
			temp1 string
			temp2 string
			temp3 string//定义三个中间量
			result float64//最终结果
		)
		line,err:=reader.ReadString('\n')//将top.txt文件按行读取
		if err == io.EOF {
			if len(line)!=0{
				fmt.Println("完成")//最后一行没有相关数据，输出完成
			}
			break
		}
		if err!=nil{
			fmt.Println("read failed")
			return
		}
		if CpuRe1.MatchString(line){//如果读取的该行有相关数据
			temp1=CpuRe1.FindString(line)
			temp2=CpuRe2.FindString(temp1)
			result,err=strconv.ParseFloat(temp2,10)//将匹配到的结果转为floa64
			if err != nil {
				fmt.Println("cpu parse failed")
				return
			}else{
				fmt.Printf("第%d个cpu数据:%.2f\n",i,result)
				i++
				_,err=CpuData.WriteString(strconv.FormatFloat(100-result,'f',-1,64)+"\n")//100-空闲的百分比即为使用的
				if err != nil {
					fmt.Println("该cpu数据写入失败")
				}
			}
		}else if MemRe1.MatchString(line){//如果读取的该行有相关数据
			temp1=MemRe1.FindString(line)
			temp2=MemRe2.FindString(temp1)
			temp3=MemRe3.FindString(temp2)
			result,err=strconv.ParseFloat(temp3,10)//将匹配结果转为float64
			if err != nil {
				fmt.Println("mem parse failed")
				return
			}else{
				fmt.Printf("第%d个mem数据：%.2f\n",j,result)
				j++
				_,err=MemData.WriteString(strconv.FormatFloat(result,'f',-1,64)+"\n")
				if err != nil {
					fmt.Println("该mem数据写入失败")
				}
			}
		}
	}
}
