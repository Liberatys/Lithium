package communication

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Request is the interface for post and get requests that can be used to communicate over http
type Request interface {
	SetParameter(key string, value string)
	SetDestination(route string)
	SetRetryCount(count int)
	SendRequest() bool
}

type GET struct {
	IP                string
	Port              string
	Url               string
	Route             string
	Parameter         map[string]string
	RequestRepetition int
}

func NewGetRequest(ip string, port string) GET {
	request := GET{
		IP:                ip,
		Port:              port,
		Parameter:         make(map[string]string),
		RequestRepetition: 3,
	}
	return request
}

func NewGetRequestOverUrl(url string) GET {
	request := GET{
		Url: url,
	}
	return request
}

func NewGetRequestWithRoute(ip string, port string, route string) GET {
	request := GET{
		IP:   ip,
		Port: port,
	}
	request.SetDestination(route)
	return request
}

func (g *GET) SetParameter(key string, value string) {
	g.Parameter[key] = value
}

func (g *GET) SetDestination(route string) {
	g.Route = route
}

func (g *GET) SetRetryCount(count int) {
	g.RequestRepetition = count
}

func (g *GET) SendRequest() (bool, string) {
	url := ""
	if g.Url != "" {
		url = g.Url
	} else {
		if len(g.Parameter) != 0 {
			parameters := ""
			for key, value := range g.Parameter {
				parameters = fmt.Sprintf("%v&%v=%v", parameters, key, value)
			}
			url = fmt.Sprintf("%v/%v/%v?%v", g.IP, g.Port, g.Route, parameters)
		} else {
			url = fmt.Sprintf("%v/%v/%v", g.IP, g.Port, g.Route)
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err.Error()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err.Error()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err.Error()
	}
	return true, string(b)
}
