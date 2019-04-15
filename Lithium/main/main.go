package main

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"net/http"
)

func main() {
	serivce := service.CreateBasicService("Message-Service", "192.168.1.1", "443", "Message")
	serivce.AddHTTTPRoute("/ping", Read)
	serivce.SpinUpHTTPServer()
}

func Read(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong"))
}
