package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreatUser(w http.ResponseWriter, r *http.Request, p *httprouter.Params) {
	io.WriteString(w, "Create user handler")
}
