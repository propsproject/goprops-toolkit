package routing

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//Route route struct
type Route struct {
	Name         string
	Method       string
	ResourcePath string
	Version      string
	NameSpace    string
	Handler      httprouter.Handle
}

func (r *Route) Use(mw httprouter.Handle) *Route {
	r.Handler = r.use(mw)
	return r
}

func (r *Route) use(mw httprouter.Handle) httprouter.Handle {
	last := r.Handler
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			mw(w, r, ps)
		}(w, r, p)

		last(w, r, p)
	}
}

func (r *Route) String() string {
	return fmt.Sprintf("Name: %v, Method: %v, Version: %v, URI: %v",
		r.Name, r.Method, r.Version, r.GetURI(),
	)
}

// GetURI ...
func (r *Route) GetURI() string {
	return fmt.Sprintf("%v/%v%v", r.NameSpace, r.Version, r.ResourcePath)
}

func NewRoute(conf map[string]string, handler httprouter.Handle) *Route {
	return &Route{
		Name:         conf["name"],
		Method:       conf["method"],
		ResourcePath: conf["resourcePath"],
		Version:      conf["version"],
		NameSpace:    conf["namespace"],
		Handler:      handler,
	}
}

// Routes ...
type Routes []*Route

func (r *Routes) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Routes\n")
	for _, route := range *r {
		buffer.WriteString(fmt.Sprintf("\t%v\n", route.String()))
	}

	return buffer.String()
}
