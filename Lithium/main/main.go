package main

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"net/http"
)

func main() {
	service := service.CreateBasicService("Message-Service", "127.0.0.1:8001", "8002", "Message")
	service.SetConfigurationLocation("configuration.txt")
	service.LoadConfigurations()
	service.InitDiscovery("127.0.0.1", "8001", 1)
	service.AddHTTTPRoute("/ping", Read)
	go service.RunDiscovery()
	service.SetSecurityModel(false)
	service.SpinUpHTTPServer()
}

func Read(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong"))
}
