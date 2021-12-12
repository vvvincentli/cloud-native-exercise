package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Write([]byte(fmt.Sprintf("i'm alive. [%s]", time.Now().Format(time.RFC3339))))
}

// error
func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(400)
	w.Write([]byte("internal server error"))
}
