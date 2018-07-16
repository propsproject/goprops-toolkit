package routing

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

type healthy struct {
	Status string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(healthy{"UP"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var healthCheck = Route {
	Name:         "Health check endpoint for server ping",
	Method:       "GET",
	ResourcePath: "/health",
	Version:      "v1",
	NameSpace:    "/check",
	Handle:  healthCheckHandler,
}