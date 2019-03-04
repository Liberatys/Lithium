package Service

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ServiceRouter struct {
	Router mux.Router
	routes []string
}

func (serviceRouter *ServiceRouter) registerRoute(path string, function func(http.ResponseWriter, *http.Request)) {
	serviceRouter.routes = append(serviceRouter.routes, path)
	serviceRouter.Router.HandleFunc(path, function)
}

func (serviceRouter *ServiceRouter) registerDefaultRoutes() {
	serviceRouter.registerRoute("/pong", Ping)
}

func Functionaly(rw http.ResponseWriter, req *http.Request) {

}

func (serviceRouter *ServiceRouter) getRoutes() []string {
	return serviceRouter.routes
}

func Ping(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Pong"))
}
