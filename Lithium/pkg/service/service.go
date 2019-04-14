package service

import (
	"fmt"
	"github.com/Liberatys/Lithium/Lithium/pkg/logging"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
	"net/http"
	"strconv"
)

type Service struct {
	Name                   string
	Location               string
	Configuration          map[string]string
	Logger                 logger.Logger
	IdentificationSequence string
	HTTPServer             server.HTTPServer
}

func CreateBasicService(Name string, Location string, Port string) Service {
	service := Service{Name: Name, Location: Location, Configuration: make(map[string]string), HTTPServer: server.InitializeBaiscHTTTPServer(Port)}
	return service
}

func (service *Service) SetLogger(logger logger.Logger) {
	service.Logger = logger
}

func (service *Service) AddConfiguration(Key string, Value string) {
	service.Configuration[Key] = Value
}

func (service *Service) GetConfigurationByKey(Key string) string {
	return service.Configuration[Key]
}

func (service *Service) AddHTTTPRoute(Route string, routeFunction func(http.ResponseWriter, *http.Request)) {
	service.HTTPServer.AddRoute(Route, routeFunction)
}

func (service *Service) GetRouteListing() []string {
	return service.HTTPServer.RouteLocationList
}

func (service *Service) SpinUpHTTPServer() {
	state := service.HTTPServer.StartHTTPServer()
	if state == false {
		service.Logger.WriteLog(fmt.Sprintf("HTTP Server was not able to start on Port: %s\nEather the port is set, or you didn't define any routes for the server", service.HTTPServer.Port))
	} else {
		service.Logger.WriteLog(fmt.Sprintf("HTTP Server now running on Port: %s with %s routes set up", service.HTTPServer.Port, strconv.Itoa(len(service.HTTPServer.RouteLocationList))))
	}
}
