package main

import (
	propslogger "github.com/propsproject/goprops-toolkit/logger"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/service"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	routeConf := map[string]string{
		"name":"Hello World",
		"resourcePath":"/helloworld",
		"method":"GET",
		"version":"v1",
		"namespace":"/check",
	}

	handler := func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Write([]byte("HELLO WORLD FROM PROPS ENGINEERING"))
	}

	route := routing.NewRoute(routeConf, handler)
	routes := []*routing.Route{route}

	logger := propslogger.NewLogger()
	name := "Example"


	helloWorld := propshttp.NewRouter(routes, 3000, name, logger).AsService()

	microservice := service.NewMicroService(name, "Example router", "v1").AddServices(helloWorld)
	microservice.Run()
}
