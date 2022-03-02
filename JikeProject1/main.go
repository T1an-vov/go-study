package main

import (
	"JikeProject1/dao"
	"JikeProject1/router"
)

func main() {
	err:=dao.SqlInit()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	r:=router.SetRouter()
	r.Run(":8080")
}
