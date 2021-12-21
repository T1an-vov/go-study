package main

import (
	"Work5/go/dao"
	"Work5/go/router"
)

func main() {
	err := dao.SqlInit()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	r := router.SetRouter()
	r.Run(":8080")
}
