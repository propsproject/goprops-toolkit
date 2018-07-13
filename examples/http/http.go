package main

import (
	"github.com/propsproject/goprops-toolkit/logger"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/propsproject/goprops-toolkit/service"
)

func main() {

	var httpService service.Service
	httpService = getRouter()

	microservice, _ := service.NewMicroService("")
	microservice.AddServices(httpService)

	microservice.Run()
}

func getRouter() *propshttp.Router {
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

	config := map[string]string{"port": "4000"}
	l := logger.NewLogger()
	name := "Example"

	return propshttp.NewRouter(routes, config, l, name)

}