# goprops-toolkit

Light wrapper packages and utilities

## Example - Microservice w/ http router

```gotemplate

package main

import (
	propslogger "github.com/propsproject/goprops-toolkit/logging"
	"github.com/propsproject/goprops-toolkit/propshttp"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/service"
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

	logger := propslogger.NewLogger("example", true)
	name := "Example"

	v := viper.New()
	v.Set("port", 3000)
	v.Set("name", "hello world")
	v.Set("version", "v1")
	helloWorld := propshttp.NewRouter(v,  routes, logger).AsService()

	microservice := service.NewMicroService(name, "Example router", "v1").AddServices(helloWorld)
	microservice.Run()
}
```
