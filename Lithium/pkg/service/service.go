package service

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/configuration"
	"github.com/Liberatys/Lithium/Lithium/pkg/database"
	"github.com/Liberatys/Lithium/Lithium/pkg/logging"
	"github.com/Liberatys/Lithium/Lithium/pkg/networking"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
	"net/http"
	"time"
)

type Service struct {
	Name                string
	Type                string
	Description         string
	Location            string
	Configuration       map[string]string
	ConfigurationPath   string
	SecurityConfig      Security
	Logger              logger.Logger
	HTTPServer          server.HTTPServer
	DatabaseConnection  database.Connection
	Discovery           networking.Discovery
	ActivationTimeStamp int64
}

func CreateBasicService(Name string, Location string, Port string, Type string) Service {
	service := Service{Name: Name, Location: Location, Configuration: make(map[string]string), HTTPServer: server.InitializeBaiscHTTTPServer(Port), Logger: logger.ConsoleLogger{}, ActivationTimeStamp: time.Now().Unix(), Type: Type,
		SecurityConfig: Security{},
	}
	return service
}

func (service *Service) InitDiscovery(IP string, Port string, Intervals int) {
	service.Discovery = networking.InitDiscovery(IP, Port, Intervals)
}

func (service *Service) RunDiscovery() {
	service.Discovery.Register(service.Configuration, service.ActivationTimeStamp, service.HTTPServer.RouteLocationList)
}

func (service *Service) SetSecurityModel(securityMode bool) {
	service.SecurityConfig.SecurityMode = securityMode
}

func (service *Service) SetConfigurationLocation(configurationFilePath string) {
	service.ConfigurationPath = configurationFilePath
}

func (service *Service) LoadConfigurations() {
	configurationContent := configuration.ReadConfigurationFile(service.ConfigurationPath)
	configurationMap := configuration.ParseGivenConfigurationFileContent(configurationContent, ":")
	for key, value := range configurationMap {
		service.Configuration[key] = value
	}
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
	if service.SecurityConfig.SecurityMode == false {
		service.HTTPServer.StartHTTPServer()
	} else {
		service.HTTPServer.StartHTTPTLSServer()
	}
}
