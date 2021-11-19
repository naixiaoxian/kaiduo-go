package main

import (
	"fmt"
	ser "kaiduo-go/http_middleware/http"
	"net/http"
)

func main() {
	fmt.Println("hello http_middleware")
	RunHttp()
}

//type content *

func RunHttp() {
	ser.Middlewares = []ser.Handler{
		login,
		loginend,
	}
	//对应的方法
	ser.Instance.AddRoute("/index2", index)
	ser.Instance.AddRoute("/", index)
	http.Handle("/", ser.Instance)
	err := http.ListenAndServe(":1000", nil)
	if err != nil {
		return
	}
}

func index(ctx *ser.Context) {
	ctx.SendJson("hello world")
}

func login(ctx *ser.Context) {
	fmt.Println("login start")
	ctx.Next()
	fmt.Println("login end")
}

func loginend(ctx *ser.Context) {
	fmt.Println("loginend start")
	ctx.Next()
	fmt.Println("loginend end")
}

//func ()  {
//
//}
