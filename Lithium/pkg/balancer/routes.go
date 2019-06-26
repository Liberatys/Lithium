/*
* The set routes a balancer should have.
* Needs to have a Register() method in order to register new services.
* Also needs a routeing method in order to route requests.
 */

package balancer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Register(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var registeredService RegisteredService
	err := decoder.Decode(&registeredService)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Invalid request"))
	} else {
		registeredService.RoutesList = strings.Split(registeredService.Routes, ";")
		registeredService.Requests = 0
		go registeredService.SpeedCheck(time.Millisecond * 1000)
		currentBalancer.Services[registeredService.Type] = append(currentBalancer.Services[registeredService.Type], registeredService)
		log.Println(fmt.Sprintf("Registration of %s from %s", registeredService.Name, registeredService.IP))
		w.Write([]byte("Registered Service"))
	}
}

func Route(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var routeRequest RouteRequest
	err := decoder.Decode(&routeRequest)
	if currentBalancer.AccessToken != "None" {
		if currentBalancer.AccessToken != routeRequest.AccessToken {
			w.Write([]byte("Invalid route-request"))
			return
		}
	}
	if err != nil {
		w.Write([]byte("Invalid route-request"))
		fmt.Println(err.Error())
	} else {
		routeRequest, success := currentBalancer.EvaluateServiceForRoute(routeRequest)
		if success == false {
			w.Write([]byte("No service found"))
			return
		}
		// TODO, this would indicate, that we can't route requests to a https site, so change the
		// protocol based on the route request.
		http.Redirect(w, req, "http://"+routeRequest.GetReRoute(), 307)
	}
}
