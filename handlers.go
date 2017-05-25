package main

import "net/http"

type theHandle struct{}

var AppHandle theHandle

func (theHandle) HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
