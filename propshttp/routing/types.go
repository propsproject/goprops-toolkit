package routing

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	defaultProtocols    = []string{"http"}
	defaultHosts        = []string{"127.0.0.1"}
	defaultStripPath    = false
	defaultPreserveHost = false
)

// Route route struct
type Route struct {
	Config  *RouteConfig
	Version string
	Service string
	Handler httprouter.Handle
}

type RouteConfig struct {
	Name         string
	Method       string
	ResourcePath string
	NameSpace    string
}

type RouteMiddleWare func(w http.ResponseWriter, r *http.Request, p httprouter.Params, next httprouter.Handle)

func (r *Route) Use(mw RouteMiddleWare) *Route {
	r.Handler = r.use(mw)
	return r
}

func (r *Route) use(mw RouteMiddleWare) httprouter.Handle {
	next := r.Handler
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		mw(w, r, p, next)
	}
}

// GetURI ...
func (r *Route) GetURI() string {
	return fmt.Sprintf("/%v%v%v", r.Version, r.Config.NameSpace, r.Config.ResourcePath)
}

// DownstreamURI uri without version prefix.
// This uri is used for kong apigateway route configuration, the version prefix will be used as the upstream path of the service.
func (r *Route) DownstreamURI() string {
	return fmt.Sprintf("/%v%v", r.Config.NameSpace, r.Config.ResourcePath)
}

func NewRoute(config *RouteConfig, handler httprouter.Handle) *Route {
	return &Route{
		Config:  config,
		Handler: handler,
	}

}

func (r *Route) String() string {
	return fmt.Sprintf("Name: %v, Method: %v, Version: %v, URI: %v",
		r.Config.Name, r.Config.Method, r.Version, r.GetURI(),
	)
}

func (r *Route) GetStrData() []string {
	return []string{r.Config.Name, r.Config.Method, r.GetURI()}
}

// Routes ...
type Routes []*Route

func (r Routes) GetStrData() [][]string {
	data := make([][]string, len(r))
	for _, route := range r {
		data = append(data, route.GetStrData())
	}

	return data
}
