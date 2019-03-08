package Balancer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (serviceMiddleWare *ServiceRegister) removeConnectionFromPool() {
	if serviceMiddleWare.CurrentConnections > 0 {
		//waiting 200mil-seconds
		//this is a prediction of how long it will take, at the slowest to process a request and send away the return.
		time.Sleep(500 * time.Millisecond)
		serviceMiddleWare.CurrentConnections -= 1
	}
}

func (serviceRegister *ServiceRegister) startSpeedTest() {
	startTime := time.Now()
	url := "http://" + serviceRegister.IP + ":" + serviceRegister.Port + "/ping"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		serviceRegister.ResponseSpeed++
	} else {
		timeout := time.Duration(1 * time.Second)
		client := &http.Client{
			Timeout: timeout,
		}
		resp, err := client.Do(req)
		if err != nil {
			serviceRegister.SpeedTestFails++
		} else {
			defer resp.Body.Close()
			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				serviceRegister.SpeedTestFails++
			} else {
				duration := time.Now().Sub(startTime)
				serviceRegister.ResponseSpeed = duration.Seconds()
				log.Println("Service: " + serviceRegister.Name + " has a latency of: " + fmt.Sprintf("%f", serviceRegister.ResponseSpeed*1000) + "ms")
				serviceRegister.SpeedTestFails = 0
				serviceRegister.ReconnectionTries = 0
			}
		}
	}
	if serviceRegister.SpeedTestFails > 0 {
		if serviceRegister.ReconnectionTries >= 2 {
			log.Println("Overwritten Service: " + serviceRegister.Name + " because we were not able to reach it")
			serviceRegister.Flagged = true
		} else {
			log.Println("SpeedTestFails for: " + serviceRegister.Name + ": " + strconv.Itoa(serviceRegister.SpeedTestFails))
			serviceRegister.ReconnectionTries++
			serviceRegister.startSpeedTest()
		}
	}
}
