package balancer

import (
	"github.com/Liberatys/Sanctuary/communication"
	"net/http"
)

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
	balancer.Services = make(map[string][]communication.SerializedService)
	setBalancer = &balancer
	return balancer
}

func (balancer *Balancer) GetService(typ string) string {
	services := balancer.Services[typ]
	if len(services) > 0 {
		return communication.Serialize(services[0])
	}
	return "No service with given type found"
}
func (balancer *Balancer) GetServices(typ string) []string {
	var services []string
	setServices := balancer.Services[typ]
	if setServices == nil {
		return nil
	}
	for _, value := range setServices {
		services = append(services, communication.Serialize(value))
	}
	return services
}

func (balancer *Balancer) Setup() {
	balancer.AddBasicRoutes()
}

func (balancer *Balancer) Start() {
	if len(balancer.HTTPServer.Routes) == 0 {
		balancer.Setup()
	}
	balancer.HTTPServer.Start()
}

func (balancer *Balancer) DecodeService(input string) communication.SerializedService {
	return communication.Decode(input)
}

func (balancer *Balancer) AddBasicRoutes() {
	balancer.HTTPServer.AddRoute("/register", Register)
	balancer.HTTPServer.AddRoute("/service", GetService)
	balancer.HTTPServer.AddRoute("/services", GetAllServices)
}

func (balancer *Balancer) AddRoute(route string, function func(w http.ResponseWriter, r *http.Request)) {
	balancer.HTTPServer.AddRoute(route, function)
}

func (balancer *Balancer) AddService(service communication.SerializedService) {
	if balancer.Services[service.Type] == nil {
		var services []communication.SerializedService
		balancer.Services[service.Type] = services
	}
	array := balancer.Services[service.Type]
	array = append(array, service)
	balancer.Services[service.Type] = array
}
