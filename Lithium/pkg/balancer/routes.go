package balancer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
		go registeredService.SpeedCheck()
		currentBalancer.Services[registeredService.Type] = append(currentBalancer.Services[registeredService.Type], registeredService)
		log.Println(fmt.Sprintf("Registration of %s from %s", registeredService.Name, registeredService.IP))
		w.Write([]byte("Registered Service"))
	}
}

func Route(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var routeRequest RouteRequest
	err := decoder.Decode(&routeRequest)
	if err != nil {
		w.Write([]byte("Invalid route-request"))
	} else {
		routeRequest, success := currentBalancer.EvaulateServiceForRoute(routeRequest)
		if success == false {
			w.Write([]byte("No service found"))
			return
		}
		http.Redirect(w, req, routeRequest.GetReRoute(), 307)
	}
}
