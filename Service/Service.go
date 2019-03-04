package Service

import (
	"net/http"
)

type ServiceInterface interface {
	GetName() string
	GetType() string
	GetDescription() string
	Ping() float64
	Alive() bool
	Response() http.Response
	Request(r http.Request)
	Reveal()
	Initialize()
	GetRoutes()
}

type Service struct {
	Name        string
	Type        string
	Description string
	IP          string
	Port        string
	Status      State
	Router      ServiceRouter
}

type State struct {
}

func (service *Service) Initialize(HomePort string, HomeIP string) {
}

func (service *Service) GetName() string {
	return service.Name
}

func (service *Service) GetType() string {
	return service.Type
}

func (service *Service) GetDescription() string {
	return service.Description
}
