package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HealthCheck(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Service alive"))
}

func Notification(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println(string(body))
	response.Write([]byte("Notification reached service"))
}
