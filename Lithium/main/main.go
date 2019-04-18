package main

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"net/http"
)

func main() {
	service := service.CreateBasicService("Message-Service", "127.0.0.1:8001", "443", "Message")
	service.SetConfigurationLocation("configuration.txt")
	service.LoadConfigurations()
	service.InitDiscovery("127.0.0.1", "8001", 1)
	service.RunDiscovery()
	service.AddHTTTPRoute("/ping", Read)
	service.SetSecurityModel(true)
	service.SpinUpHTTPServer()
}

/**
service := service.CreateBasicService("Message-Service", "192.168.1.1", "443", "Message")
	service.AddHTTTPRoute("/ping", Read)
	service.SetSecurityModel(true)
	service.SpinUpHTTPServer()
*/

func Read(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong"))
}
