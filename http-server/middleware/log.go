package middleware

import (
	"net"
	"net/http"
	"strings"
)

// response observer
type ResponseObserver struct {
	http.ResponseWriter
	Status      int
	Written     int64
	WroteHeader bool
}

// get client ip
func GetClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// Write ,write
func (o *ResponseObserver) Write(p []byte) (n int, err error) {
	if !o.WroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.Written += int64(n)
	return
}

// WriteHeader ,write code to http header
func (o *ResponseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.WroteHeader {
		return
	}
	o.WroteHeader = true
	o.Status = code
}


