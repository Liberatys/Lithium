package networking

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
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

func (discovery *Discovery) Register(configuration map[string]string, Timestamp int64, RouteLocationList []string) bool {
	/**
	The required information for the balancer that is needed in order to perform the discovery is:
		- Service-Type
		- Service-Name
		- Service-IP
		- Service-Port
		- Route-List
		- Activation-Timestamp
	*/
	percent, _ := cpu.Percent(time.Second, true)
	serviceConfiguration := make(map[string]string)
	serviceConfiguration["type"] = configuration["type"]
	serviceConfiguration["name"] = configuration["name"]
	serviceConfiguration["ip"] = GetOutboundIP().String()
	serviceConfiguration["port"] = configuration["port"]
	serviceConfiguration["load"] = fmt.Sprintf("%f", percent)
	serviceConfiguration["activation"] = strconv.FormatInt(Timestamp, 10)
	serviceConfiguration["protocol"] = configuration["protocol"]
	routeListing := ""
	for _, element := range RouteLocationList {
		routeListing += element + ";"
	}
	serviceConfiguration["routes"] = routeListing
	connectionSequence := fmt.Sprintf("https://%s:%s/register", discovery.DiscoveryIP, discovery.DiscoveryPort)
	responseBody, error := SendPOSTRequest(connectionSequence, serviceConfiguration)
	fmt.Println(responseBody)
	if error == false {
		return error
	}
	/**
	Store response parsed as DiscoveryReponse for later use
	*/
	var discoveryResult DiscoveryResponse
	discovery.DiscoveryResults = append(discovery.DiscoveryResults, discoveryResult)
	json.Unmarshal([]byte(responseBody), discoveryResult)
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
