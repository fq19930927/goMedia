package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

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
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
