package routers

import (
	"cloud-native-exercise/http-server/controller"
	"net/http"
)

func RouterBinding() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", controller.Healthz)
	return mux
}
