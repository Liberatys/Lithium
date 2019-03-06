package Balancer

import (
	"encoding/json"
	"github.com/Liberatys/Lithium/Service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
	The ServiceMiddleWare is recording the repsonseTime of the service and is also implementing a method for easy logging of the service.
*/

type BalancerRouter struct {
	Router    mux.Router
	Services  []ServiceMiddleWare
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
	/*
		Parse the sent information that is used to redirect the user.
	 */
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
			//sends the client to the service that is handling these types of requests.
			http.Redirect(rw, request, "http://"+serviceInformation+""+destination.Destination, 301)
		}
	}
}

func DiscoverService(rw http.ResponseWriter, r *http.Request) {
	var serviceToRegister Service.Service
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &serviceToRegister)
	newServiceMiddleWare := ServiceMiddleWare{Service: serviceToRegister, Flagged: true}
	go newServiceMiddleWare.startSpeedTest()
	balacerRouter.Services = append(balacerRouter.Services, newServiceMiddleWare)
	rw.Write([]byte("Registered Service"))
}

/*
	This method runs over the services, and searches for a service that matches the request, after that
	it is looking for the fastest service of the given ones.
	If a service is found, it will return the IP+""+Port
	Else it will return "None Found"


	Service speed is termined by the last response-Time * the current Connections.
	1 second after we redirected the service, we will decrement the current Connection counter.
	This is a prediction and is just a temporary system.
*/

func (balancerRouter *BalancerRouter) findServicePointForType(serviceType string) string {
	var fastestService ServiceMiddleWare
	var fastestWeightedSpeed float64
	for _, element := range balancerRouter.Services {
		if strings.Compare(strings.Trim(element.Service.Type, " "), strings.Trim(serviceType, " ")) == 0 {
			if fastestService.Flagged == false {
				fastestService = element
				fastestWeightedSpeed = calculateWeightedSpeed(element)
			} else {
				elementValue := calculateWeightedSpeed(element)
				if elementValue < fastestWeightedSpeed && element.Service.Flagged == false {
					fastestWeightedSpeed = elementValue
					fastestService = element
				}
			}
		}
	}
	if fastestService.Flagged != false {
		fastestService.CurrentConnections += 1
		go fastestService.removeConnectionFromPool()
		return fastestService.Service.IP + ":" + fastestService.Service.Port
	}
	return "None Found"
}

func calculateWeightedSpeed(element ServiceMiddleWare) float64 {
	elementValue := element.ResponseSpeed
	if element.CurrentConnections > 0 {
		elementValue *= float64(element.CurrentConnections)
	}
	return elementValue
}
