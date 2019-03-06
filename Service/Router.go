package Service

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ServiceRouter struct {
	Router  mux.Router
	routes  []string
	Blocked bool
}

func (serviceRouter *ServiceRouter) InitServiceRouter() {
	serviceRouter.Router = mux.Router{}
	serviceRouter.Blocked = true
	serviceRouter.registerDefaultRoutes()
}

func (serviceRouter *ServiceRouter) OpenConnections() {
	serviceRouter.Blocked = false
}

func (serviceRouter *ServiceRouter) registerRoute(path string, function func(http.ResponseWriter, *http.Request)) {
	serviceRouter.routes = append(serviceRouter.routes, path)
	serviceRouter.Router.HandleFunc(path, function)
}

func (serviceRouter *ServiceRouter) registerDefaultRoutes() {
	serviceRouter.registerRoute("/ping", Ping)
}

func (serviceRouter *ServiceRouter) getRoutes() []string {
	return serviceRouter.routes
}

func Ping(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte("Pong"))
}
