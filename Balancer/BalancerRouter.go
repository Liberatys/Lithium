package Balancer

import (
	"encoding/json"
	"github.com/Liberatys/Lithium/Service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
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

type Destination struct {
	Destination     string
	Destinationtype string
}

func APIGateWay(rw http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	decoder := json.NewDecoder(request.Body)
	var destination Destination
	err := decoder.Decode(&destination)
	if err != nil {
		rw.Write([]byte("Please provide destination and type"))
	} else {
		serviceInformation := balacerRouter.findServicePointForType(destination.Destinationtype)
		if serviceInformation == "None Found" {
			rw.Write([]byte("No Service found for this service type"))
		} else {



			http.Redirect(rw, request, "http://"+serviceInformation+""+destination.Destination, 301)
		}
	}
}

func DiscoverService(rw http.ResponseWriter, r *http.Request) {
	var serviceToRegister Service.Service
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &serviceToRegister)
	balacerRouter.Services = append(balacerRouter.Services, serviceToRegister)
	rw.Write([]byte("Registered Service"))
}

func (balancerRouter *BalancerRouter) RetreaveServices() []Service.Service {
	return balancerRouter.Services
}

func (balancerRouter *BalancerRouter) findServicePointForType(serviceType string) string {
	for _, element := range balancerRouter.Services {
		if strings.Compare(strings.Trim(element.Type, " "), strings.Trim(serviceType, " ")) == 0 {
			return element.IP + ":" + element.Port
		}
	}
	return "None Found"
}
