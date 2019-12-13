package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreatUser)
	return router
}

func main() {
	r := RegisterHandles()
	http.ListenAndServe(":8000", r)
}
