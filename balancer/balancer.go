package balancer

import "github.com/Liberatys/Sanctuary/communication"

type Balancer struct {
	Port       string
	Name       string
	HTTPServer communication.HTTPConnection
	Services   map[string][]communication.SerializedService
}

func NewBalancer(name string, port string) Balancer {
	balancer := Balancer{
		Name:       name,
		Port:       port,
		HTTPServer: communication.NewHTTPServer("127.0.0.1", port),
	}
	return balancer
}

func (balancer *Balancer) Setup() {
	balancer.AddBasicRoutes()
}

func (balancer *Balancer) Start() {

}

func DecodeService(input string) communication.SerializedService {
	return communication.Decode(input)
}

func (balancer *Balancer) AddBasicRoutes() {

}

func (balancer *Balancer) AddRoute() {

}
