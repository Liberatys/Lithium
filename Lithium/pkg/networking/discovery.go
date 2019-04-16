package networking

import (
	"encoding/json"
	"fmt"
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"strconv"
)

type DiscoveryRoutine interface {
	Ping() bool
	Register() bool
	LastResponse() DiscoveryResponse
}
type Discovery struct {
	DiscoveryIP   string
	DiscoveryPort string
	//Currently Intervals is not doing anything, but in the future, it will be used to make fitness checks to the balancer.
	DiscoveryIntervals int
	DiscoveryFailure   bool
	DiscoveryResults   []DiscoveryResponse
}

func InitDiscovery(IP string, Port string, Intervals int) Discovery {
	discovery := Discovery{DiscoveryIP: IP, DiscoveryPort: Port, DiscoveryIntervals: Intervals}
	return discovery
}

type DiscoveryResponse struct {
	Status         string
	Identification string
	Command        string
}

func (discovery *Discovery) Register(service *service.Service) bool {
	/**
	The required information for the balancer that is needed in order to perform the discovery is:
		- Service-Type
		- Service-Name
		- Service-IP
		- Service-Port
		- Route-List
		- Activation-Timestamp
	*/
	serviceConfiguration := make(map[string]string)
	serviceConfiguration["type"] = service.Configuration["type"]
	serviceConfiguration["name"] = service.Configuration["name"]
	serviceConfiguration["ip"] = GetOutboundIP().String()
	serviceConfiguration["port"] = service.HTTPServer.Port
	serviceConfiguration["activation"] = strconv.FormatInt(service.ActivationTimeStamp, 10)
	routeListing := ""
	for _, element := range service.HTTPServer.RouteLocationList {
		routeListing += element + ";"
	}
	serviceConfiguration["routes"] = routeListing
	service.Configuration = serviceConfiguration
	connectionSequence := fmt.Sprintf("%s:%s", discovery.DiscoveryIP, discovery.DiscoveryPort)
	responseBody, error := SendPOSTRequest(connectionSequence, serviceConfiguration)
	if error == false {
		return error
	}
	/**
	Store response parsed as DiscoveryReponse for later use
	*/
	var discoveryResult DiscoveryResponse
	json.Unmarshal([]byte(responseBody), discoveryResult)
	discovery.DiscoveryResults = append(discovery.DiscoveryResults, discoveryResult)
	return true
}

func (discovery *Discovery) Ping() bool {
	connectionSequence := fmt.Sprintf("%s:%s", discovery.DiscoveryIP, discovery.DiscoveryPort)
	_, success := SendGETRequest(connectionSequence)
	return success
}

/**

Return the last response in the slice stored for Discovery
*/
func (discovery *Discovery) LastResponse() DiscoveryResponse {
	return discovery.DiscoveryResults[len(discovery.DiscoveryResults)-1]
}
