package Discovery

import "github.com/Liberatys/Lithium/Service"

type Discoverer struct {
	HomeIP      string
	HomePort    string
	HomePayload string
	Service     Service.Service
}

func (discoverer *Discoverer) InitDiscoverer(HomePort string, HomeIP string, ServiceType string) {
	discoverer.HomeIP = HomeIP
	discoverer.HomePort = HomePort

}

func (discoverer *Discoverer) RunDiscovery() {

}
