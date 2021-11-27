package main

import (
	"cloud-native-exercise/http-server/internal/configs"
	"cloud-native-exercise/http-server/middleware"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	defer shutdown()

	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := ""
	flag.StringVar(&configFile, "app-config", "", "application config file path.")
	flag.Parse()
	if configFile == "" {
		configFile = os.Getenv("app-config")
	}

	log.Println(fmt.Sprintf("LoadConfig %s", configFile))
	if err := configs.LoadConfig(configFile); err != nil {
		log.Println(fmt.Sprintf("loadConfig %s failed, error:%s", configFile, err))
		os.Exit(-1)
	}
	c := configs.GetConfig()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	hostServer(c.App.Host, c.App.Port, errs)
	//f, ferr := os.Create("/tmp/http-server")
	//if ferr != nil {
	//	f.WriteString("http-server started.")
	//}
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
