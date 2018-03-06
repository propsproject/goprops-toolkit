package http

import (
	"fmt"
	"net/http"
	"strconv"

	lgr "github.com/propsproject/go-utils/logger/v2"
	"github.com/propsproject/props-ws-service/http/routing"
	"github.com/rs/cors"
	goji "goji.io"
	"goji.io/pat"
)

//Router ...
type Router struct {
	mux    *goji.Mux
	addr   string
	routes []routing.Route
	port   int
	logger *lgr.LoggerWrapper
}

//NewRouter returns a new router
func NewRouter(routes []routing.Route, config map[string]string) *Router {
	port, _ := strconv.Atoi(config["port"])
	addr := fmt.Sprintf(":%s", config["port"])
	router := &Router{
		mux:    goji.NewMux(),
		addr:   addr,
		routes: append(routes, routing.DefaultRoutes...),
		port:   port,
		logger: lgr.Logger,
	}

	//TODO: setup proper CORS configuration
	router.mux.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)

	router.registerRoutes()
	return router
}

func (r *Router) registerRoutes() {
	for _, route := range r.routes {
		switch route.Method {
		case "GET":
			r.mux.HandleFunc(pat.Get(route.Pattern), route.HandlerFunc)
		case "POST":
			r.mux.HandleFunc(pat.Post(route.Pattern), route.HandlerFunc)
		case "PUT":
			r.mux.HandleFunc(pat.Put(route.Pattern), route.HandlerFunc)
		case "DELETE":
			r.mux.HandleFunc(pat.Delete(route.Pattern), route.HandlerFunc)
		default:
			err := fmt.Errorf("Unsupported method type (%v) on route (%v), supported methods are GET POST PUT DELETE UPDATE", route.Method, route.Name)
			panic(err)
		}
	}
}

func (r *Router) logRoutes() {
	for _, route := range r.routes {
		r.logger.Info(route.Name,
			lgr.Field{Key: "method", Value: route.Method},
			lgr.Field{Key: "pattern", Value: route.Pattern})
	}
}

//Start start http router
func (r *Router) Start() {
	r.logger.Info("Starting HTTP server ", lgr.Field{Key: "port", Value: strconv.Itoa(r.port)})
	r.logRoutes()
	http.ListenAndServe(r.addr, r.mux)
}
