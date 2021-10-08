package controller

import (
	"net/http"
	"os"
)

// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	w.Header().Set("VERSION",os.Getenv("VERSION"))
	w.Write([]byte("i'm alive."))
}

// error
func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION",os.Getenv("VERSION"))
	w.WriteHeader(400)
	w.Write([]byte("internal server error"))
}