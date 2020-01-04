package communication

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
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
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("certs"),
	}
	timeOutInSeconds := time.Second * 4
	ser := &http.Server{
		Addr:         ":" + server.Port,
		ReadTimeout:  timeOutInSeconds,
		WriteTimeout: timeOutInSeconds,
		IdleTimeout:  timeOutInSeconds,
		Handler:      &CORSRouterDecorator{router},
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":3400", certManager.GetCertificate(nil))
	err := ser.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(fmt.Sprintf("Was not able to spin up http server on port: %v because: %v", server.Port, err.Error()))
	}
}
