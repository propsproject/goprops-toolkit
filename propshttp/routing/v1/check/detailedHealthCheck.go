package check

import (
	"encoding/json"
	"github.com/thoas/stats"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/propsproject/goprops-toolkit/propshttp/routing"

)

type DetailedHealthCheckResponse struct {
	Stats            *stats.Data `json:"stats"`
	MicroserviceName string      `json:"microservice_name"`
	Type             string      `json:"type"`
}

var detailedHealthCheckConf = map[string]string{
	"name":"Detailed Health check endpoint for server ping",
	"resourcePath":"/detailed-health",
	"method":"GET",
	"namespace":namespace,
	"version":version,
}

var detailedHealthCheckHandler = func(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := DetailedHealthCheckResponse{
		Stats:            stats.New().Data(),
		MicroserviceName: "TODO",
		Type:             "TODO",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var detailedHealthCheck = routing.NewRoute(detailedHealthCheckConf, detailedHealthCheckHandler)
