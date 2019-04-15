package service

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/database"
	"github.com/Liberatys/Lithium/Lithium/pkg/logging"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
	"net/http"
	"time"
)

type Service struct {
	Name                   string
	Type                   string
	Location               string
	Configuration          map[string]string
	Logger                 logger.Logger
	IdentificationSequence string
	HTTPServer             server.HTTPServer
	DatabaseConnection     database.Connection
	ActivationTimeStamp    int64
}

func CreateBasicService(Name string, Location string, Port string, Type string) Service {
	service := Service{Name: Name, Location: Location, Configuration: make(map[string]string), HTTPServer: server.InitializeBaiscHTTTPServer(Port), Logger: logger.ConsoleLogger{}, ActivationTimeStamp: time.Now().Unix(), Type: Type}
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
	service.HTTPServer.StartHTTPServer()

}
