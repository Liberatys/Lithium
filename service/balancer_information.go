package service

import (
	"fmt"

	"github.com/Liberatys/Sanctuary/communication"
)

type ServiceBalancer struct {
	IP         string
	Port       string
	Region     string
	SecureKey  string
	Configured bool
}

func NewServiceBalancer(ip string, port string, region string) ServiceBalancer {
	serviceBalancer := ServiceBalancer{
		IP:         ip,
		Port:       port,
		Region:     region,
		Configured: true,
	}
	return serviceBalancer
}

func (serviceBalancer *ServiceBalancer) Connect(service *Service) (error, string) {
	post := communication.NewPostRequest(serviceBalancer.IP, serviceBalancer.Port, "", "/register", service.Serialize())
	return post.SendRequest()
}

func (serviceBalancer *ServiceBalancer) HealthCheck() (bool, string) {
	getRequest := communication.NewGetRequestOverUrl(fmt.Sprintf("%v:%v/%v", serviceBalancer.IP, serviceBalancer.Port, "health"))
	//replace with the breaker method to call the server
	err, message := getRequest.SendRequest()
	if err != nil {
		return false, message
	}
	return true, message
}

func (serviceBalancer *ServiceBalancer) AddSecureKey(key string) {
	serviceBalancer.SecureKey = key
}
