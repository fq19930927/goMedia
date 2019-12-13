package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreatUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create user handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
