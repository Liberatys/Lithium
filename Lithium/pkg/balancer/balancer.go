package balancer

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/configuration"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
)

type Balancer struct {
	Configuration map[string]string
	Port          string
	IP            string
	Services      map[string][]RegisteredService
	HTTPServer    server.HTTPServer
}

var currentBalancer Balancer

func CreateNewBalancer(Port string, IP string) Balancer {
	balancer := Balancer{Port: Port, IP: IP, HTTPServer: server.InitializeBaiscHTTTPServer(Port), Configuration: make(map[string]string), Services: make(map[string][]RegisteredService)}
	currentBalancer = balancer
	return balancer
}

func (balancer *Balancer) RegisterBasicRoutes() {
	balancer.HTTPServer.AddRoute("/register", Register)
	balancer.HTTPServer.AddRoute("/route", Route)
}

func (balancer *Balancer) SpinUpHTTP() {
	balancer.HTTPServer.StartHTTPServer()
}

func (balancer *Balancer) LoadConfigurations(FileLocation string) {
	configurationContent := configuration.ReadConfigurationFile(FileLocation)
	configurationMap := configuration.ParseGivenConfigurationFileContent(configurationContent, ":")
	for key, value := range configurationMap {
		balancer.Configuration[key] = value
	}
}

func (balancer *Balancer) EvaulateServiceForRoute(routeRequest RouteRequest) (RouteRequest, bool) {
	if val, ok := balancer.Services[routeRequest.Type]; ok {
		lowestValue := 0
		lowestValueElement := RegisteredService{}
		for _, value := range val {
			if lowestValue == 0 {
				lowestValue = value.Requests
				lowestValueElement = value
			} else {
				if lowestValue > value.Requests {
					lowestValue = value.Requests
					lowestValueElement = value
				}
			}
		}
		lowestValueElement.IncrementRequests()
		routeRequest.IP = lowestValueElement.IP
		routeRequest.Port = lowestValueElement.Port
		return routeRequest, true
	} else {
		return RouteRequest{}, false
	}
}
