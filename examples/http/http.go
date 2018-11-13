package main

import (
	"github.com/propsproject/goprops-toolkit/logging"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

func main() {
	handler := func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		writer.Write([]byte("HELLO WORLD FROM PROPS ENGINEERING"))
	}

	route := routing.NewRoute(&routing.RouteConfig{Name: "Hello World", ResourcePath: "/helloworld", Method: "GET", NameSpace: "/check"}, handler)
	routes := []*routing.Route{route}
	name := "Example"

	logger := logging.Get(name, true)

	viper.Set("port", 3000)
	viper.Set("name", "hello world")
	viper.Set("version", "v1")

	helloWorld := propshttp.NewRouter(viper.GetViper(), logger, routes)

	helloWorld.Start()
}
