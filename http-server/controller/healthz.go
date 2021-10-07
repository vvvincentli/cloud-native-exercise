package controller

import "net/http"

// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	w.Header().Set("server-header","testtttttt")
	w.Write([]byte("i'm alive."))
}
