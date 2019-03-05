package Balancer

import (
	"github.com/Liberatys/Lithium/Service"
	"github.com/gorilla/mux"
	"net/http"
)

type BalancerRouter struct {
	Router    mux.Router
	Services  []Service.Service
	Initiated bool
}

var balacerRouter BalancerRouter

func InitNewRouter() BalancerRouter {
	newBalancerRouter := BalancerRouter{Router: mux.Router{}, Initiated: true}
	newBalancerRouter.registerDefaultRoutes()
	balacerRouter = newBalancerRouter
	return newBalancerRouter
}

func (balancerRouter *BalancerRouter) AddRoute(routePath string, methodToRegister func(http.ResponseWriter, *http.Request)) {
	balancerRouter.Router.HandleFunc(routePath, methodToRegister)
}

func (balancerRouter *BalancerRouter) registerDefaultRoutes() {
	balancerRouter.Router.HandleFunc("/apigateway", APIGateWay)
}

func APIGateWay(rw http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	destinationURL := request.FormValue("destination")
	serviceRequestType := request.FormValue("destinationtype")
	serviceInfomration := balacerRouter.findServicePointForType(serviceRequestType)
	if serviceInfomration == "None Found" {
		rw.Write([]byte("No Service found for this service type"))
	} else {
		http.Redirect(rw, request, "http://"+serviceInfomration+""+destinationURL, 301)
	}
}

func (balancerRouter *BalancerRouter) RetreaveServices() []Service.Service {
	return balancerRouter.Services
}

func (balancerRouter *BalancerRouter) findServicePointForType(serviceType string) string {
	for _, element := range balancerRouter.Services {
		if element.Type == serviceType {
			return element.IP + ":" + element.Port
		}
	}
	return "None Found"
}
