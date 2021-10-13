package main

import (
	"cloud-native-exercise/http-server/middleware"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//	runtime.GOMAXPROCS(runtime.NumCPU())

	defer shutdown()
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	hostServer("", "8088", errs)
	fmt.Println(fmt.Sprintf("exit: %v", <-errs))
}

// host server
func hostServer(addr, port string, errs chan error) {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: middleware.RouterBinding(),
	}
	go func(server http.Server, err chan error) {
		fmt.Println("server started.")
		err <- server.ListenAndServe()
	}(server, errs)

}

//shutdown , clear resource
func shutdown() {

}
