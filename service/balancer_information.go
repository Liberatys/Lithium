package service

type ServiceBalancer struct {
	IP         string
	Port       string
	Region     string
	SecureKey  string
	Configured bool
}

func newServiceBalancer(ip string, port string, region string) ServiceBalancer {
	serviceBalancer := ServiceBalancer{
		IP:         ip,
		Port:       port,
		Region:     region,
		Configured: true,
	}
	return serviceBalancer
}

func (serviceBalancer *ServiceBalancer) AddSecureKey(key string) {
	serviceBalancer.SecureKey = key
}