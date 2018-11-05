package check

import (
	"net/http"
	"encoding/json"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/julienschmidt/httprouter"
)

type healthy struct {
	Status string `json:"status"`
}

var healthCheckConf = &routing.RouteConfig{
	Name:         "Health check - endpoint for server ping",
	ResourcePath: "/health",
	Method:       "GET",
	NameSpace:    namespace,
}

var healthCheckHandler = func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(healthy{"UP"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var healthCheck = routing.NewRoute(healthCheckConf, healthCheckHandler)