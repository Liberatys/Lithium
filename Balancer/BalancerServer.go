package Balancer

import (
	"log"
	"net/http"
	"time"
)

/**
The SpinUpServer is creating a local http-server for the balancer and is setting the router as the Handler.
This structure makes it possible, that we can plugin routes and handlers on a very high level.
*/
func (balancer *Balancer) SpinUpServer() {
	timeOutInSeconds := time.Second * 1
	server := &http.Server{
		Addr:         ":" + balancer.Port,
		ReadTimeout:  timeOutInSeconds,
		WriteTimeout: timeOutInSeconds,
		IdleTimeout:  timeOutInSeconds,
		Handler:      &CORSRouterDecorator{&balancer.Router.Router},
	}
	log.Println("Now listening on Port: " + balancer.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Was not able to create the Balancer")
	}
}
