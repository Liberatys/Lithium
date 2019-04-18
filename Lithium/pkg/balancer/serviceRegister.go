package balancer

import (
	"github.com/Liberatys/Lithium/Lithium/pkg/networking"
	"log"
	"time"
)

type RegisteredService struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	Activation   string `json:"activation"`
	Routes       string `json:"routes"`
	Protocol     string `json:"protocol"`
	Requests     int
	RoutesList   []string
	RequestDelay int64
}

func (registeredService *RegisteredService) SpeedCheck() {
	log.Println("Speedchecking for Service")
	ip := registeredService.IP
	if registeredService.IP == networking.GetOutboundIP().String() {
		ip = "127.0.0.1"
	}
	start := time.Now()
	log.Println(networking.SendGETRequest(registeredService.Protocol + "://" + ip + ":" + registeredService.Port + "/ping"))
	elapsed := time.Since(start)
	registeredService.RequestDelay = elapsed.Nanoseconds()
}

func (registrationService *RegisteredService) IncrementRequests() {
	registrationService.Requests++
}
