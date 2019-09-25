package balancer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var setBalancer *Balancer

func Register(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	service := setBalancer.DecodeService(string(body))
	setBalancer.AddService(service)
	w.Write([]byte(fmt.Sprintf("The service %v has been registered to %v:%v", service.Name, setBalancer.Name, setBalancer.Port)))
	return
}

func GetService(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	responseMessage := setBalancer.GetService(string(body))
	w.Write([]byte(responseMessage))
}
