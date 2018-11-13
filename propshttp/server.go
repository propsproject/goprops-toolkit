package propshttp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/olekukonko/tablewriter"
	"github.com/propsproject/goprops-toolkit/logging"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/propsproject/goprops-toolkit/propshttp/routing/v1/check"
	"github.com/spf13/viper"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
)

//Router ...
type Router struct {
	mux     *httprouter.Router
	addr    string
	port    int
	Name    string
	Version string
	Type    string
	server  *http.Server
	routes  routing.Routes
	logger  *logging.Logger
	Config  *viper.Viper
}

func (r *Router) ServiceURL() string {
	return fmt.Sprintf("%s/%s", r.Config.GetString("url"), r.Config.GetString("version"))
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
	r.logger.WithField("port", strconv.Itoa(r.port)).Info(fmt.Sprintf("Starting HTTP Router %s\n", r.Name))
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
func (r *Router) Start() {
	r.logRoutes()
	idleConnsClosed := make(chan struct{})
	go r.listen(idleConnsClosed)

	if err := r.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		r.logger.Fatalf("error starting listener for http server")
	}

	<-idleConnsClosed
}

func (r *Router) listen(idleConnsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigint

	// We received an interrupt signal, shut down.
	r.logger.Warnf("Attempting graceful shutdown of HTTP server: %s", r.Name)

	if err := r.server.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		r.logger.Warnf("Could not gracefully shutdown HTTP server: %s", r.Name)
	}
	close(idleConnsClosed)
}

func (r *Router) Port() string {
	return string(r.port)
}

func (r *Router) WithLogger(logger *logging.Logger) *Router {
	r.logger = logger
	return r
}

//NewRouter returns a new router
func NewRouter(v *viper.Viper, logger *logging.Logger, routes routing.Routes) *Router {
	port := v.GetInt("port")
	name := v.GetString("name")
	addr := fmt.Sprintf(":%v", port)

	mux := httprouter.New(httprouter.WithServiceName(name))

	check.DetailedHealthCheckName = name
	check.DetailedHealthCheckType = "http"

	router := &Router{
		mux:     mux,
		addr:    addr,
		server:  &http.Server{Addr: addr, Handler: mux},
		routes:  append(routes, check.DefaultRoutes...),
		port:    port,
		logger:  logger,
		Version: v.GetString("version"),
		Name:    name,
		Config:  v,
	}

	//TODO: setup proper CORS configuration
	router.registerRoutes()
	return router
}
