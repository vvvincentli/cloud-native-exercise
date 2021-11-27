package middleware

import (
	"cloud-native-exercise/http-server/pkg/controller"
	"fmt"
	"net/http"
	"time"
)

// router binding
func RouterBinding() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", withLogging(controller.Healthz))
	mux.HandleFunc("/error", withLogging(controller.Error))
	return mux
}

// write log
func withLogging(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		o := &ResponseObserver{ResponseWriter: writer}
		h(o, request)
		fmt.Println(fmt.Sprintf("%s - - [%s] %q %d %d %q %q",
			GetClientIP(request),
			time.Now().Format(time.RFC3339Nano),
			fmt.Sprintf("%s %s %s", request.Method, request.URL, request.Proto),
			o.Status,
			o.Written,
			request.Referer(),
			request.UserAgent()))
	}
}
