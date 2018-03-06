package routing

import (
	"encoding/json"
	"net/http"

	"github.com/propsproject/props-ws-service/instruments"
)

type healthy struct {
	Status string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	span := instruments.StartRootSpan("HTTP/Ping")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(healthy{"UP"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//healthCheck a route for this handler to be registered by server
var healthCheck = Route{
	Name:        "Health check endpoint for server ping",
	Method:      "GET",
	Pattern:     "/healthcheck",
	HandlerFunc: healthCheckHandler,
}

//DefaultRoutes default routes for your server
var DefaultRoutes = []Route{healthCheck}
