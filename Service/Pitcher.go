package Service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/**
	The PitchServer method is the first method that a service runs after it has been created.
	This method will call the server with all its information and it will tell the server where to find the service and how to call it.
	Also with this, we get a central server to oversee the status of all servers. On this server we can handle the load and sort it to the services.
*/

func (service *Service) PitchServer() {
	payLoad, err := json.Marshal(service)
	log.Println("Sending discovery to server")
	if err != nil {
	} else {
		url := "http://" + service.HomePoint.HomeIP + ":" + service.HomePoint.HomePort + "" + service.HomePoint.HomeEndPoint
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payLoad))
		if err != nil {
			log.Println(err.Error())
			log.Fatal("Error at creating Request for discovery")
		}
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
		} else {
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err.Error())
				log.Fatal("Error while executing Discover-Reguest")
			}
			defer resp.Body.Close()
			if err != nil {
			} else {
				if resp.StatusCode > 400 && resp.StatusCode < 511 {
					log.Fatal("Was not able to call the server, shutting down the Service")
				} else {
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Fatal("Was not able to connect to the server")
					} else {
						bodySequence := string(body[:])
						if bodySequence == "Registered Service" {
							log.Println("Service has been registered")
						}
					}
				}
			}
		}
	}
}
