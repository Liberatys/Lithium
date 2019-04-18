package balancer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var registeredService RegisteredService
	err := decoder.Decode(&registeredService)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Invalid request"))
	}
	log.Println(fmt.Sprintf("Registration of %s from %s", registeredService.Name, registeredService.IP))
}
