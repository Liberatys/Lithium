package service

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/database"
	"github.com/Liberatys/Lithium/Lithium/pkg/logging"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
	"net/http"
	"time"
)

type Service struct {
	Name              string
	Type              string
	Description       string
	Location          string
	Configuration     map[string]string
	ConfigurationPath string
	//SecurityMode is a variable for the state of tls or standart http. 1 Is TLS and 0 is http.
	SecurityMode           bool
	Logger                 logger.Logger
	IdentificationSequence string
	HTTPServer             server.HTTPServer
	DatabaseConnection     database.Connection
	ActivationTimeStamp    int64
}

func CreateBasicService(Name string, Location string, Port string, Type string) Service {
	service := Service{Name: Name, Location: Location, Configuration: make(map[string]string), HTTPServer: server.InitializeBaiscHTTTPServer(Port), Logger: logger.ConsoleLogger{}, ActivationTimeStamp: time.Now().Unix(), Type: Type, SecurityMode: false,
		IdentificationSequence: "Not Identified",
	}
	return service
}

func (service *Service) SetSecurityModel(securityMode bool) {
	service.SecurityMode = securityMode
}

func (service *Service) SetConfigurationLocation(configurationFilePath string) {
	service.ConfigurationPath = configurationFilePath
}

func (service *Service) LoadConfigurations() {
	configurationContent := ReadConfigurationFile(service.ConfigurationPath)
	configurationMap := ParseGivenConfigurationFileContent(configurationContent, ":")
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
	if service.SecurityMode == false {
		service.HTTPServer.StartHTTPServer()
	} else {
		service.HTTPServer.StartHTTPTLSServer()
	}
}
