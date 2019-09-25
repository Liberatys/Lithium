package main

import "github.com/Liberatys/Sanctuary/balancer"

func main() {
	balancer := balancer.NewBalancer("Test", "3400")
	balancer.Setup()
	balancer.Start()
}
