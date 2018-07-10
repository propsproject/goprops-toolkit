package main

import (
	"context"
	"github.com/propsproject/goprops-toolkit/logger"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"net/http"
	"os"
	"syscall"
)

func main() {

	routes := routing.Routes{
		routing.Route{
			Name:         "HelloUniverse",
			Method:       "GET",
			ResourcePath: "/helloworld",
			Version:      "v1",
			NameSpace:    "",
			HandlerFunc: func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("HELLO UNIVERSE"))
			},
		},
	}

	config := map[string]string{"port": "3000"}
	l := logger.NewLogger()
	name := "Example"
	ctx, cancel := context.WithCancel(context.Background())

	router := propshttp.NewRouter(routes, config, l, name, ctx)
	go router.Start()

	sig1 := make(chan os.Signal, 1)
	sig2 := make(chan os.Signal, syscall.SIGINT)
	for {
		select {
		case <-sig1:
			l.Info("Stopping microservice signal 1")
			cancel()
			os.Exit(0)
		case <-sig2:
			l.Info("Stopping microservice signal 2")
			cancel()
			os.Exit(0)
		}
	}
}
