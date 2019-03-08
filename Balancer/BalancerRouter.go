package Balancer

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

/**
	The ServiceMiddleWare is recording the repsonseTime of the service and is also implementing a method for easy logging of the service.
*/

var serviceMap = make(map[string][]ServiceRegister)

type BalancerRouter struct {
	Router             mux.Router
	Services           []ServiceRegister
	Initiated          bool
	UsingLithiumLogger bool
}

var balacerRouter BalancerRouter

func InitNewRouter(useLithiumLogger bool) BalancerRouter {
	newBalancerRouter := BalancerRouter{Router: mux.Router{}, Initiated: true, UsingLithiumLogger: useLithiumLogger}
	balacerRouter = newBalancerRouter
	return newBalancerRouter
}

type ServiceRegister struct {
	Name               string
	IP                 string
	Port               string
	Type               string
	Flagged            bool
	ResponseSpeed      float64
	CurrentConnections int
	SpeedTestFails     int
	ReconnectionTries  int
}

type Destination struct {
	Destination     string
	Destinationtype string
}

func APIGateWay(rw http.ResponseWriter, request *http.Request) {
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
			http.Redirect(rw, request, "http://"+serviceInformation+""+destination.Destination, 307)
		}
	}
}

func DiscoverService(rw http.ResponseWriter, r *http.Request) {
	var serviceToRegister ServiceRegister
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &serviceToRegister)
	go serviceToRegister.startSpeedTest()
	if _, ok := serviceMap[serviceToRegister.Type]; ok {
	} else {
		serviceMap[serviceToRegister.Type] = []ServiceRegister{}
	}
	serviceMap[serviceToRegister.Type] = append(serviceMap[serviceToRegister.Type], serviceToRegister)
	balacerRouter.Services = append(balacerRouter.Services, serviceToRegister)
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
	var fastestService ServiceRegister
	fastestService.Flagged = true
	var fastestWeightedSpeed float64
	if _, ok := serviceMap[serviceType]; ok {
		for _, element := range serviceMap[serviceType] {
			if element.Flagged == false {
				if fastestService.Flagged == true {
					fastestService = element
				} else {
					currentElementSpeed := calculateWeightedSpeed(element)
					if currentElementSpeed < fastestWeightedSpeed {
						fastestService = element
						fastestWeightedSpeed = currentElementSpeed
					}
				}
			}
		}
	} else {
		return "None Found"
	}
	if fastestService.Flagged == false {
		fastestService.CurrentConnections += 1
		go fastestService.removeConnectionFromPool()
		return fastestService.IP + ":" + fastestService.Port
	}
	return "None Found"
}

func calculateWeightedSpeed(element ServiceRegister) float64 {
	elementValue := element.ResponseSpeed
	if element.CurrentConnections > 0 {
		elementValue *= float64(element.CurrentConnections)
	}
	return elementValue
}
