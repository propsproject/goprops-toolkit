package check

import "github.com/propsproject/goprops-toolkit/propshttp/routing"

//DefaultRoutes default routes for your server
var DefaultRoutes = routing.Routes{healthCheck, detailedHealthCheck}
