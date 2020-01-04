package communication

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPSConnection struct {
	IP        string
	Port      string
	Routes    map[string]func(w http.ResponseWriter, r *http.Request)
	Activated bool
	Cert      string
	Key       string
}

func NewHTTPSServer(ip string, port string) *HTTPSConnection {
	server := HTTPSConnection{
		IP:        ip,
		Port:      port,
		Routes:    make(map[string]func(w http.ResponseWriter, r *http.Request)),
		Activated: true,
	}
	return &server
}

func (server *HTTPSConnection) SetState(state bool) {
	server.Activated = state
}

func (server *HTTPSConnection) GetState() bool {
	return server.Activated
}

func (server *HTTPSConnection) AddRoute(route string, function func(w http.ResponseWriter, r *http.Request)) {
	server.Routes[route] = function
}

func (server *HTTPSConnection) Start() {
	router := mux.NewRouter().StrictSlash(true)
	for key, value := range server.Routes {
		router.HandleFunc(key, value)
	}
	timeOutInSeconds := time.Second * 4
	ser := &http.Server{
		Addr:         ":" + server.Port,
		ReadTimeout:  timeOutInSeconds,
		WriteTimeout: timeOutInSeconds,
		IdleTimeout:  timeOutInSeconds,
		Handler:      &CORSRouterDecorator{router},
	}
	err := ser.ListenAndServeTLS(server.Cert, server.Key)
	if err != nil {
		log.Fatal(fmt.Sprintf("Was not able to spin up http server on port: %v because: %v", server.Port, err.Error()))
	}
}
