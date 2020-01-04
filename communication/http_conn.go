package communication

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPConnection struct {
	IP        string
	Port      string
	Routes    map[string]func(w http.ResponseWriter, r *http.Request)
	Activated bool
}

func NewHTTPServer(ip string, port string) *HTTPConnection {
	server := HTTPConnection{
		IP:        ip,
		Port:      port,
		Routes:    make(map[string]func(w http.ResponseWriter, r *http.Request)),
		Activated: true,
	}
	return &server
}

func (server *HTTPConnection) AddRoute(route string, function func(w http.ResponseWriter, r *http.Request)) {
	server.Routes[route] = function
}

func (server *HTTPConnection) SetState(state bool) {
	server.Activated = state
}

func (server *HTTPConnection) GetState() bool {
	return server.Activated
}
func (server *HTTPConnection) Start() {
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
	fmt.Println(fmt.Sprintf("HTTP-Server spun up on port :%v", server.Port))
	err := ser.ListenAndServe()
	if err != nil {
		log.Fatal(fmt.Sprintf("Was not able to spin up http server on port: %v because: %v", server.Port, err.Error()))
	}
}
