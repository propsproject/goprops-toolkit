package main

import (
	"github.com/propsproject/goprops-toolkit/logger"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"net/http"
	"github.com/propsproject/goprops-toolkit/utils"
	"os"
	"github.com/julienschmidt/httprouter"
)

func main() {

	routes := routing.Routes{
		routing.Route{
			Name:         "HelloUniverse",
			Method:       "GET",
			ResourcePath: "/helloworld",
			Version:      "v1",
			NameSpace:    "",
			HandlerFunc: func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
				writer.Write([]byte("HELLO UNIVERSE"))
			},
		},
	}

	config := map[string]string{"port": "3000"}
	l := logger.NewLogger()
	name := "Example"

	router := propshttp.NewRouter(routes, config, l, name)
	go router.Start()

	utils.NewGracefulShutDownListener([]os.Signal{os.Kill, os.Interrupt}).Listen(router.ShutdownSig)
}


