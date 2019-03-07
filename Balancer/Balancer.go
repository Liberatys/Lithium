package Balancer

import (
	"net/http"
	"time"
)

type Balancer struct {
	CreationTimeStamp int64
	Router            BalancerRouter
	Port              string
}

//CreateNewBalancer is setting the needed variables and is returning an instance for usage.
func CreateNewBalancer(Port string, useLithiumLogger bool) Balancer {
	newBalancer := Balancer{CreationTimeStamp: time.Now().Unix(), Port: Port}
	newBalancer.initBalancer(useLithiumLogger)
	newBalancer.registerDefaultRoutes()
	return newBalancer
}

func (balancer *Balancer) initBalancer(useLitiumLogger bool) {
	balancer.Router = InitNewRouter(useLitiumLogger)
}

func (balancer *Balancer) AddRoute(routePath string, methodToRegister func(http.ResponseWriter, *http.Request)) {
	balancer.Router.Router.HandleFunc(routePath, methodToRegister)
}

func (balancer *Balancer) registerDefaultRoutes() {
	balancer.AddRoute("/apigateway", APIGateWay)
	balancer.AddRoute("/serviceDiscovery", DiscoverService)
	if balacerRouter.UsingLithiumLogger == true {
		balancer.AddRoute("/lithium/Logger", LogServiceAction)
	}
}
