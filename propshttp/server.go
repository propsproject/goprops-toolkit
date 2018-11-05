package propshttp

import (
	"fmt"
	"net/http"
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/propsproject/goprops-toolkit/logging"
	"github.com/propsproject/goprops-toolkit/propshttp/kong"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/propshttp/routing/v1/check"
	"github.com/propsproject/goprops-toolkit/service"
	"github.com/propsproject/goprops-toolkit/utils/sharedconf"
	"github.com/spf13/viper"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
	"os"
	"os/signal"
	"strconv"
)

var (
	defaultProtocol       = "http"
	defaultConnectTimeout = 10000
	defaultWriteTimeout   = 10000
	defaultReadTimeout    = 10000
	defaultRetries        = 3
)

//Router ...
type Router struct {
	mux         *httprouter.Router
	addr        string
	port        int
	Name        string
	Version     string
	Type        string
	server      *http.Server
	routes      routing.Routes
	logger      *logging.PLogger
	shutdownSig chan bool
	Config      *viper.Viper
}

func (r *Router) ServiceURL() string {
	return fmt.Sprintf("%s/%s", r.Config.GetString("url"), r.Config.GetString("version"))
}

func (r *Router) AsService() service.Service {
	return r
}

func (r *Router) registerRoutes() {
	for _, route := range r.routes {
		route.Version = r.Version
		route.Service = r.Name
		switch route.Config.Method {
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
			err := fmt.Errorf("unsupported method type (%v) on route (%v), supported methods are GET POST PUT DELETE UPDATE", route.Config.Method, route.Config.Name)
			panic(err)
		}
	}
}

func (r *Router) logRoutes() {
	r.logger.Info(fmt.Sprintf("Starting HTTP Router %s", r.Name), logging.Field{Key: "port", Value: strconv.Itoa(r.port)}, "\n").Log()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Method", "URI"})
	table.SetFooter([]string{
		fmt.Sprintf("router name: %s", r.Name),
		fmt.Sprintf("version: %s", r.Version),
		fmt.Sprintf("port: %v", r.port)})

	table.SetBorder(false)
	table.SetCaption(true, fmt.Sprintf("Routes registered to HTTP Server: %s", r.Name))

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor})

	table.SetFooterColor(
		tablewriter.Colors{tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.FgMagentaColor})

	table.SetColumnSeparator("|")
	table.SetRowSeparator("_")
	table.SetAutoWrapText(true)

	for _, route := range r.routes {
		table.Append(route.GetStrData())
	}

	table.Render()
}

//Start start http router
func (r *Router) Start(regCh chan sharedconf.ConsulRegistration) {
	r.logRoutes()

	idleConnsClosed := make(chan struct{})
	go r.listen(idleConnsClosed)

	if err := r.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		r.logger.Fatal("error starting or shutting down HTTP server", r.Name).Log()
	}

	if regCh != nil {
		regCh <- sharedconf.ConsulRegistration{Name: "http", Port: r.port}
	}

	<-idleConnsClosed

	r.shutdownSig <- true
}

func (r *Router) listen(idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// We received an interrupt signal, shut down.
	r.logger.Warn(fmt.Sprintf("Attempting graceful shutdown of HTTP server: %s", r.Name)).Log()

	if err := r.server.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		r.logger.Warn(fmt.Sprintf("Could not gracefully shutdown HTTP server: %s", r.Name)).Log()
	}
	close(idleConnsClosed)
}

func (r *Router) ShutDownSig() chan bool {
	return r.shutdownSig
}

func (r *Router) Port() string {
	return string(r.port)
}

//NewRouter returns a new router
func NewRouter(v *viper.Viper, routes routing.Routes, logger *logging.PLogger) *Router {

	port := v.GetInt("port")
	name := v.GetString("name")
	addr := fmt.Sprintf(":%v", port)

	mux := httprouter.New(httprouter.WithServiceName(name))

	check.DetailedHealthCheckName = name
	check.DetailedHealthCheckType = "http"

	router := &Router {
		mux:         mux,
		addr:        addr,
		server:      &http.Server{Addr: addr, Handler: mux},
		routes:      append(routes, check.DefaultRoutes...),
		port:        port,
		Version:     v.GetString("version"),
		logger:      logger,
		shutdownSig: make(chan bool),
		Name:        name,
		Config:      v,
	}

	//TODO: setup proper CORS configuration
	router.registerRoutes()
	return router.SetDefaults()
}
