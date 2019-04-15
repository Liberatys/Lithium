package main

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"net/http"
)

func main() {
	service := service.CreateBasicService("Message-Service", "192.168.1.1", "443", "Message")
	service.AddHTTTPRoute("/ping", Read)
	service.SpinUpHTTPServer()
}

func Read(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong"))
}
