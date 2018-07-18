package propshttp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/logger"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"github.com/propsproject/goprops-toolkit/propshttp/routing/v1/check"
	"github.com/propsproject/goprops-toolkit/service"
	"sync"
	"context"
)

//Router ...
type Router struct {
	mux         *httprouter.Router
	addr        string
	server      *http.Server
	routes      routing.Routes
	port        int
	logger      *logger.Wrapper
	shutdownSig chan *sync.WaitGroup
	Name        string
}

//NewRouter returns a new router
func NewRouter(routes routing.Routes, config map[string]string, logger *logger.Wrapper, name string) *Router {
	port, _ := strconv.Atoi(config["port"])
	addr := fmt.Sprintf(":%s", config["port"])
	mux := httprouter.New(httprouter.WithServiceName(config["name"]))
	router := &Router {
		mux:         mux,
		addr:        addr,
		server:      &http.Server{Addr: addr, Handler: mux},
		routes:      append(routes, check.DefaultRoutes...),
		port:        port,
		logger:      logger,
		shutdownSig: make(chan *sync.WaitGroup),
		Name:        name,
	}

	//TODO: setup proper CORS configuration
	router.registerRoutes()
	return router
}

func (r *Router) AsService() service.Service {
	return r
}

func (r *Router) registerRoutes() {
	for _, route := range r.routes {
		switch route.Method {
		case "GET":
			r.mux.Handle("GET", route.GetURI(), route.Handler)
		case "POST":
			r.mux.POST(route.GetURI(), route.Handler)
		case "PUT":
			r.mux.PUT(route.GetURI(), route.Handler)
		case "DELETE":
			r.mux.DELETE(route.GetURI(), route.Handler)
		case "OPTIONS":
			r.mux.OPTIONS(route.GetURI(), route.Handler)
		default:
			err := fmt.Errorf("unsupported method type (%v) on route (%v), supported methods are GET POST PUT DELETE UPDATE", route.Method, route.Name)
			panic(err)
		}
	}
}

func (r *Router) logRoutes() {
	r.logger.Info(r.routes.String())
}

//Start start http router
func (r *Router) Start(regCh chan sharedconf.ConsulRegistration) {
	r.logger.Info(fmt.Sprintf("Starting HTTP Router %s", r.Name), logger.Field{Key: "port", Value: strconv.Itoa(r.port)})
	r.logRoutes()

	go func() {
		if err := r.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	if regCh != nil {
		regCh <- sharedconf.ConsulRegistration{Name: "http", Port: r.port}
	}
	r.WaitForShutdown()
}

func (r *Router) WaitForShutdown()  {
	for {
		if wg := <- r.shutdownSig; wg != nil {
			r.logger.Warn(fmt.Sprintf("Attempting graceful shutdown of HTTP server: %s", r.Name))
			err := r.server.Shutdown(context.Background())
			if err != nil {
				r.logger.Warn(fmt.Sprintf("Could not gracefully shutdown HTTP server: %s", r.Name))
			}
			wg.Done()
			break
		}
	}
}

func (r *Router) ShutDownSig() chan *sync.WaitGroup {
	return r.shutdownSig
}

func (r *Router) Port() string {
	return string(r.port)
}