package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	//创建用户
	router.POST("/user", CreatUser)
	//用户登录
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandles()
	http.ListenAndServe(":8000", r)
}
