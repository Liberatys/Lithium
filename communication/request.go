package communication

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Request is the interface for post and get requests that can be used to communicate over http
type Request interface {
	SetParameter(key string, value string)
	SetDestination(route string)
	SetRetryCount(count int)
	SendRequest() (error, string)
}

type POST struct {
	IP    string
	Port  string
	Url   string
	Route string
	Body  string
}

func NewPostRequest(ip, port, url, route, body string) POST {
	req := POST{
		IP:    ip,
		Port:  port,
		Url:   url,
		Route: route,
		Body:  body,
	}
	return req
}

func (post *POST) SendRequest() (error, string) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%v:%v/%v", post.IP, post.Port, post.Route), bytes.NewBuffer([]byte(post.Body)))
	if err != nil {
		return err, err.Error()
	}
	req.Header.Set("X-Custom-Header", "sanctuary")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, err.Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return nil, string(body)
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

func (g *GET) SendRequest() (error, string) {
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
		return err, err.Error()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, err.Error()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, err.Error()
	}
	return nil, string(b)
}
