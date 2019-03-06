package Balancer

import (
	"fmt"
	"github.com/Liberatys/Lithium/Service"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ServiceMiddleWare struct {
	Service            Service.Service
	ResponseSpeed      float64
	SpeedTestFails     int
	Flagged            bool
	ReconnectionTries  int
	CurrentConnections int
}

func (serviceMiddleWare *ServiceMiddleWare) removeConnectionFromPool() {
	if serviceMiddleWare.CurrentConnections > 0 {
		//waiting 200mil-seconds
		//this is a prediction of how long it will take, at the slowest to process a request and send away the return.
		time.Sleep(500 * time.Millisecond)
		serviceMiddleWare.CurrentConnections -= 1
	}
}

func (serviceMiddleWare *ServiceMiddleWare) startSpeedTest() {
	startTime := time.Now()
	url := "http://" + serviceMiddleWare.Service.IP + ":" + serviceMiddleWare.Service.Port + "/ping"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		serviceMiddleWare.ResponseSpeed++
	} else {
		timeout := time.Duration(1 * time.Second)
		client := &http.Client{
			Timeout: timeout,
		}
		resp, err := client.Do(req)
		if err != nil {
			serviceMiddleWare.SpeedTestFails++
		} else {
			defer resp.Body.Close()
			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				serviceMiddleWare.SpeedTestFails++
			} else {
				duration := time.Now().Sub(startTime)
				serviceMiddleWare.ResponseSpeed = duration.Seconds()
				log.Println("Service: " + serviceMiddleWare.Service.Name + " has a latency of: " + fmt.Sprintf("%f", serviceMiddleWare.ResponseSpeed*1000) + "ms")
				serviceMiddleWare.SpeedTestFails = 0
				serviceMiddleWare.ReconnectionTries = 0
			}
		}
	}
	if serviceMiddleWare.SpeedTestFails > 0 {
		if serviceMiddleWare.ReconnectionTries >= 2 {
			log.Println("Overwritten Service: " + serviceMiddleWare.Service.Name + " because we were not able to reach it")
			serviceMiddleWare.Service.Flagged = true
		} else {
			log.Println("SpeedTestFails for: " + serviceMiddleWare.Service.Name + ": " + strconv.Itoa(serviceMiddleWare.SpeedTestFails))
			serviceMiddleWare.ReconnectionTries++
			serviceMiddleWare.startSpeedTest()
		}
	}
}
