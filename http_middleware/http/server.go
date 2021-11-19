package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler func(ctx *Context)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Handlers []Handler
	index    int8
}

type Server struct {
	route map[string]Handler
}

func (ser *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//继承的这个方法是重点
	ctx := &Context{
		Request:  r,
		Response: w,
		index:    -1,
	}
	ctx.Handlers = append(Middlewares, func(ctx *Context) {
		ser.handler(ctx)
	})
	ctx.Next()
}

var Middlewares []Handler
var Instance *Server

func init() {
	Instance = new(Server)
	Instance.route = make(map[string]Handler)
	fmt.Println("serverInstance===> init success")
}

// NewServer
//department
func NewServer() *Server {
	return &Server{
		route: make(map[string]Handler),
	}
}

func (ser *Server) handler(ctx *Context) {
	url := ctx.Request.URL.Path
	if h := ser.match(url); h != nil {
		h(ctx)
	} else {
		fmt.Println("http ==> 404 找不到url", url)
		ctx.SendJson("http ==> 404 找不到url " + url)
	}
}

func (ser *Server) match(url string) Handler {
	if h, ok := ser.route[url]; ok {
		return h
	}
	return nil
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.Handlers)) {
		ctx.Handlers[ctx.index](ctx)
		ctx.index++
	}
}

func (ser *Server) AddRoute(urlPath string, handler Handler) {
	ser.route[urlPath] = handler
}

func (ctx *Context) SendJson(v interface{}) {
	p, _ := json.Marshal(v)
	_, err := ctx.Response.Write(p)
	if err != nil {
		return
	}
}
