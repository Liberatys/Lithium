package service

import (
	"fmt"
	"github.com/Liberatys/Sanctuary/communication"
	"net/http"
)

//Service structure for the base implementation of a service that can be expanded with other "components"
type Service struct {
	Name        string
	Type        string
	Description string
	IP          string
	Port        string
	Balancer    ServiceBalancer
	HTTPServer  communication.HTTPConnection
}

func newService(name string, typ string, description string) Service {
	service := Service{
		Name:        name,
		Type:        typ,
		Description: description,
		Balancer:    ServiceBalancer{},
		Port:        "",
		HTTPServer:  nil,
	}
	return service
}

func (service *Service) SetServiceBalancer(serviceBalancer ServiceBalancer) {
	service.Balancer = serviceBalancer
}

func (service *Service) ActivateHTTPServer(port string) {
	service.Port = port
	service.HTTPServer = communication.NewHTTPServer(service.IP, service.Port)
}

func (service *Service) AddHTTPRoute(route string, fn func(w http.ResponseWriter, r *http.Request)) (bool, string) {
	if service.HTTPServer.Activated {
		service.HTTPServer.AddRoute(route, fn)
		return true, fmt.Sprintf("Route %v added", route)
	}
	return false, "HTTPServer is not activated ... \n Activate using service.ActivateHTTPServer(port)"
}
