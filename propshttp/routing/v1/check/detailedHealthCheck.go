package check

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"
	"github.com/thoas/stats"
	"net/http"
)

var (
	DetailedHealthCheckName string
	DetailedHealthCheckType string
)

type DetailedHealthCheckResponse struct {
	Stats *stats.Data `json:"stats"`
	Name  string      `json:"http_router_name"`
	Type  string      `json:"type"`
}

var detailedHealthCheckConf = map[string]string{
	"name":         "Detailed Health check endpoint for server ping",
	"resourcePath": "/detailed-health",
	"method":       "GET",
	"namespace":    namespace,
	"version":      version,
}

var detailedHealthCheckHandler = func(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := DetailedHealthCheckResponse{
		Stats: stats.New().Data(),
		Name:  DetailedHealthCheckName,
		Type:  DetailedHealthCheckType,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var detailedHealthCheck = routing.NewRoute(detailedHealthCheckConf, detailedHealthCheckHandler)
