package Balancer

import (
	"log"
	"net/http"
	"time"
)

func (balancer *Balancer) SpinUpServer() {
	timeOutInSeconds := time.Second * 1
	server := &http.Server{
		Addr:         ":" + balancer.Port,
		ReadTimeout:  timeOutInSeconds,
		WriteTimeout: timeOutInSeconds,
		IdleTimeout:  timeOutInSeconds,
		Handler:      &CORSRouterDecorator{&balancer.Router.Router},
	}
	log.Println("Not listening on Port: " + balancer.Port)
	server.ListenAndServe()
}
