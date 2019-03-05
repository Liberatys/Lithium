package Balancer

import (
	"time"
)

type Balancer struct {
	CreationTimeStamp int64
	Router            BalancerRouter
	Port              string
}

func CreateNewBalancer() Balancer {
	newBalancer := Balancer{CreationTimeStamp: time.Now().Unix()}
	return newBalancer
}

func (balancer *Balancer) InitBalancer() {
	balancer.Router = InitNewRouter()

}
