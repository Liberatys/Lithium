package balancer

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/configuration"
	"github.com/Liberatys/Lithium/Lithium/pkg/server"
)

type Balancer struct {
	Configuration map[string]string
	Port          string
	IP            string
	Services      []RegisteredService
	HTTPServer    server.HTTPServer
}

func CreateNewBalancer(Port string, IP string) Balancer {
	balancer := Balancer{Port: Port, IP: IP, HTTPServer: server.InitializeBaiscHTTTPServer(Port), Configuration: make(map[string]string)}
	return balancer
}

func (balancer *Balancer) RegisterBasicRoutes() {
	balancer.HTTPServer.AddRoute("/register", Register)
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
