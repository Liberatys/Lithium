package balancer

type RegisteredService struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	IP         string `json:"ip"`
	Port       string `json:"port"`
	Activation string `json:"activation"`
	Routes     string `json:"routes"`
}
