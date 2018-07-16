package propshttp

import (
	"fmt"
	"net/http"
	"strconv"

	"context"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/logger"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
	"time"
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
)

//Router ...
type Router struct {
	mux         *httprouter.Router
	addr        string
	server      *http.Server
	routes      routing.Routes
	port        int
	logger      *logger.Wrapper
	shutdownSig chan bool
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
		routes:      append(routes, routing.DefaultRoutes...),
		port:        port,
		logger:      logger,
		shutdownSig: make(chan bool),
		Name:        name,
	}

	//TODO: setup proper CORS configuration
	router.mux.HandleOPTIONS = true
	router.registerRoutes()
	return router
}

func (r *Router) registerRoutes() {
	for _, route := range r.routes {
		switch route.Method {
		case "GET":
			r.mux.GET(route.GetURI(), route.HandlerFunc)
		case "POST":
			r.mux.POST(route.GetURI(), route.HandlerFunc)
		case "PUT":
			r.mux.PUT(route.GetURI(), route.HandlerFunc)
		case "DELETE":
			r.mux.DELETE(route.GetURI(), route.HandlerFunc)
		case "OPTIONS":
			r.mux.OPTIONS(route.GetURI(), route.HandlerFunc)
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
	r.logger.Info("Starting HTTP server ", logger.Field{Key: "port", Value: strconv.Itoa(r.port)})
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
		select {
		case <-r.shutdownSig:
			r.logger.Warn(fmt.Sprintf("Attempting graceful shutdown of HTTP server: %s", r.Name))
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			r.server.Shutdown(ctx)
			cancel()
			r.shutdownSig <- false
		}
	}
}

func (r *Router) ShutDownSig() chan bool {
	return r.shutdownSig
}

func (r *Router) Port() string {
	return string(r.port)
}