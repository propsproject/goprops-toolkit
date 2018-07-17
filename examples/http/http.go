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
		writer.Write([]byte("HELLO UNIVERSE"))
	}

	route := routing.NewRoute(routeConf, handler)
	routes := []routing.Route{route}

	routerServiceConf := map[string]string{"port": "3000"}
	logger := propslogger.NewLogger()
	name := "Example"


	helloWorld := propshttp.NewRouter(routes, routerServiceConf, logger, name).AsService()

	microservice := service.NewMicroService().AddServices(helloWorld)
	microservice.Run()
}
