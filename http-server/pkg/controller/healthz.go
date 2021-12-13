package controller

import (
	"cloud-native-exercise/http-server/middleware/metrics"
	"cloud-native-exercise/middleware/metrics"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	tm := metrics.NewTimer()
	defer tm.ObserveTotal()
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	time.Sleep(time.Duration(intn(50, 2000) * int(time.Microsecond)))
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Write([]byte(fmt.Sprintf("i'm alive. [%s]", time.Now().Format(time.RFC3339))))
}

func intn(min, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Intn(max-min)
}

// error
func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(400)
	w.Write([]byte("internal server error"))
}
