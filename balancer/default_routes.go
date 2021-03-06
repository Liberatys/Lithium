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

func GetAllServices(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	messages := setBalancer.GetServices(string(body))
	if len(messages) < 1 {
		w.Write([]byte("No servies found in this type"))
		return
	}
	for _, value := range messages {
		w.Write([]byte(value))
		w.Write([]byte("\n"))
	}
	return
}

// TODO: implement the relay system

// relay a message to all servies in a network
func RelayCommand(w http.ResponseWriter, r *http.Request) {

}

// relay a message to a group of serviess
func RelayCommandToRec(w http.ResponseWriter, r *http.Request) {

}

//TODO: implement methods for service update

// TODO: Add a validation method for the default routes so they can only be triggered by requests with a valid "ID"
