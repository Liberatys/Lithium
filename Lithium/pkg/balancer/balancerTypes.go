package balancer

type RouteRequest struct {
	Type       string `json:"type"`
	Endpoint   string `json:"Endpoint"`
	Importance int    `json:"importance"`
	IP         string
	Port       string
}

func (routeRequest *RouteRequest) GetReRoute() string {
	routeLocation := routeRequest.IP + ":" + routeRequest.Port + "/" + routeRequest.Endpoint
	return routeLocation
}
