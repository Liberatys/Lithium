package Service

import (
	"github.com/Liberatys/Lithium/Discovery"
	"net/http"
)

type ServiceInterface interface {
	Name()
	Type()
	Ping() float64
	Alive() bool
	Response() http.Response
	Request(r http.Request)
	Reveal()
	Initialize()
}

type Service struct {
	Name        string
	Type        string
	Description string
	IP          string
	Port        string
	Status      State
	Discoverer  Discovery.Discoverer
}

type State struct {
}

func (service *Service) Initialize() {

}
