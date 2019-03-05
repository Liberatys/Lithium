package Balancer

import (
	"encoding/json"
	"github.com/Liberatys/Lithium/Service"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	balacerRouter = newBalancerRouter
	return newBalancerRouter
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

func DiscoverService(rw http.ResponseWriter, r *http.Request) {
	var serviceToRegister Service.Service
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &serviceToRegister)
	balacerRouter.Services = append(balacerRouter.Services, serviceToRegister)
	rw.Write([]byte("Registered Service: " + string(serviceToRegister.Name)))
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
