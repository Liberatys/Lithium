package balancer

import (
	"net/http"

	"github.com/Liberatys/Sanctuary/communication"
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

/*
*
* Must be tested and should also be modifiable by the user of the framework, because different usage is in need of different filtering.
* TODO: Implementation of an api for modifying this part of the code.
*
 */
func (balancer *Balancer) GetService(typ string) string {
	services := balancer.Services[typ]
	if len(services) <= 0 {
		return "No service with given type found"
	}
	lowest_index := services[0].LoadIndex
	opt_service := services[0]
	for _, value := range services {
		if value.LoadIndex < lowest_index {
			lowest_index = value.LoadIndex
			opt_service = value
		}
	}
	//TODO: fix this issue, wrong return type.... serialize it and return as string
	return opt_service
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

// TODO: have the default route summed into a component that then can be added to the balancer
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
