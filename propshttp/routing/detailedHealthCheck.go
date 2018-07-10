package routing

import (
	"encoding/json"
	"github.com/thoas/stats"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type DetailedHealthCheckResponse struct {
	Stats            *stats.Data `json:"stats"`
	MicroserviceName string     `json:"microservice_name"`
	Type             string     `json:"type"`
}

func detailedHealthCheckHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := DetailedHealthCheckResponse{
		Stats: stats.New().Data(),
		MicroserviceName: "TODO",
		Type: "TODO",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var detailedHealthCheck = Route {
	Name:         "Detailed Health check endpoint for server ping",
	NameSpace:    "/check",
	Version:      "v1",
	ResourcePath: "/detailed-health",
	Method:       "GET",
	HandlerFunc:  detailedHealthCheckHandler,
}
