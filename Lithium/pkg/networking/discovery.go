package networking

import (
	"encoding/json"
	"fmt"
	"github.com/Liberatys/Lithium/Lithium/pkg/service"
	"strconv"
)

type DiscoveryRoutine interface {
	Ping()
	Register()
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

func (discovery *Discovery) Register(service *service.Service) {
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
	serviceConfiguration["type"] = service.Type
	serviceConfiguration["name"] = service.Name
	serviceConfiguration["ip"] = GetOutboundIP().String()
	serviceConfiguration["port"] = service.HTTPServer.Port
	routeListing := ""
	for _, element := range service.HTTPServer.RouteLocationList {
		routeListing += element + ";"
	}
	serviceConfiguration["routes"] = routeListing
	serviceConfiguration["activation"] = strconv.FormatInt(service.ActivationTimeStamp, 10)
	service.Configuration = serviceConfiguration
	connectionSequence := fmt.Sprintf("%s:%s", discovery.DiscoveryIP, discovery.DiscoveryPort)
	responseBody := SendPOSTRequest(connectionSequence, serviceConfiguration)
	/**
	Store response parsed as DiscoveryReponse for later use
	*/
	var discoveryResult DiscoveryResponse
	json.Unmarshal([]byte(responseBody), discoveryResult)
	discovery.DiscoveryResults = append(discovery.DiscoveryResults, discoveryResult)
}

/**
Return the last response in the slice stored for Discovery
*/
func (discovery *Discovery) LastResponse() DiscoveryResponse {
	return discovery.DiscoveryResults[len(discovery.DiscoveryResults)-1]
}
