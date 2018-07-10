package http

import (
	"fmt"
	"net/http"
	"strconv"

	"context"
	"github.com/propsproject/goprops-toolkit/http/routing"
	"github.com/propsproject/goprops-toolkit/logger"
	"github.com/rs/cors"
	"goji.io"
	"goji.io/pat"
	"time"
)

func ContextWithLogger(l logger.Wrapper) context.Context {
	return context.WithValue(context.Background(), "logger", l)
}

//Router ...
type Router struct {
	mux     *goji.Mux
	addr    string
	server  *http.Server
	routes  routing.Routes
	port    int
	logger  *logger.Wrapper
	Context context.Context
	Name    string
}

//NewRouter returns a new router
func NewRouter(routes routing.Routes, config map[string]string, logger *logger.Wrapper, name string, ctx context.Context) *Router {
	port, _ := strconv.Atoi(config["port"])
	addr := fmt.Sprintf(":%s", config["port"])
	mux := goji.NewMux()
	router := &Router {
		mux:     mux,
		addr:    addr,
		server:  &http.Server{Addr: addr, Handler: mux},
		routes:  append(routes, routing.DefaultRoutes...),
		port:    port,
		logger:  logger,
		Context: ctx,
		Name:    name,
	}

	//TODO: setup proper CORS configuration
	router.mux.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}).Handler)

	router.registerRoutes()
	return router
}

func (r *Router) registerRoutes() {
	for _, route := range r.routes {
		switch route.Method {
		case "GET":
			r.mux.HandleFunc(pat.Get(route.GetURI()), route.HandlerFunc)
		case "POST":
			r.mux.HandleFunc(pat.Post(route.GetURI()), route.HandlerFunc)
		case "PUT":
			r.mux.HandleFunc(pat.Put(route.GetURI()), route.HandlerFunc)
		case "DELETE":
			r.mux.HandleFunc(pat.Delete(route.GetURI()), route.HandlerFunc)
		case "OPTIONS":
			r.mux.HandleFunc(pat.Options(route.GetURI()), route.HandlerFunc)
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
func (r *Router) Start() {
	r.logger.Info("Starting HTTP server ", logger.Field{Key: "port", Value: strconv.Itoa(r.port)})
	r.logRoutes()

	go func() {
		if err := r.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	<-r.Context.Done()
	r.logger.Warn(fmt.Sprintf("Attempting graceful shutdown of HTTP server: %s", r.Name))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r.server.Shutdown(ctx)
}
