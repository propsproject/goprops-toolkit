package routing

import (
	"encoding/json"
	"net/http"
)

type healthy struct {
	Status string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(healthy{"UP"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var healthCheck = Route{
	Name:         "Health check endpoint for server ping",
	Method:       "GET",
	ResourcePath: "/health",
	Version:      "v1",
	NameSpace:    "/check",
	HandlerFunc:  healthCheckHandler,
}

//DefaultRoutes default routes for your server
var DefaultRoutes = Routes{healthCheck}
