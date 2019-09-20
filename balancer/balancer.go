package balancer

import "github.com/Liberatys/Sanctuary/communication"

type Balancer struct {
	Port       string
	Name       string
	HTTPServer communication.HTTPConnection
}


