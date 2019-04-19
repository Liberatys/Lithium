package balancer

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/networking"
	"strings"
)

type RouteRequest struct {
	Type       string `json:"type"`
	Endpoint   string `json:"endpoint"`
	Importance string `json:"importance"`
	IP         string
	Port       string
}

func (routeRequest *RouteRequest) GetReRoute() string {
	IP := routeRequest.IP
	if strings.Trim(IP, " ") == strings.Trim(networking.GetOutboundIP().String(), " ") {
		IP = "127.0.0.1"
	}
	routeLocation := IP + ":" + routeRequest.Port + routeRequest.Endpoint
	return routeLocation
}
