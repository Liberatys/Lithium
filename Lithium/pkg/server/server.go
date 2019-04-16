package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kabukky/httpscerts"
)

type HTTPServer struct {
	RouteLocationList []string
	Router            *mux.Router
	HTTPServer        *http.Server
	Port              string
}

func InitializeBaiscHTTTPServer(Port string) HTTPServer {
	httpServer := HTTPServer{Router: mux.NewRouter().StrictSlash(true), Port: Port}
	//setting a low repsonse time for server and client, should solve the problem of DDOS with blocking like slow loris.
	timeOut := time.Second * 1
	server := &http.Server{
		Addr:         ":" + Port,
		ReadTimeout:  timeOut,
		WriteTimeout: timeOut,
		IdleTimeout:  timeOut,
		//warp all requests into CORS, in order to allow set information to be passed.
		Handler: &CORSRouterDecorator{httpServer.Router},
	}
	httpServer.HTTPServer = server
	return httpServer
}

func (httpServer *HTTPServer) AddRoute(RouteLocation string, RouteFunction func(w http.ResponseWriter, req *http.Request)) {
	httpServer.RouteLocationList = append(httpServer.RouteLocationList, RouteLocation)
	httpServer.Router.HandleFunc(RouteLocation, RouteFunction)
}

func (httpServer *HTTPServer) StartHTTPTLSServer() bool {
	if len(httpServer.RouteLocationList) == 0 {
		return false
	}
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", fmt.Sprintf("127.0.0.1:%s", httpServer.Port))
		if err != nil {
			log.Fatal("Was not able to generate required keys")
		}
	}
	err = httpServer.HTTPServer.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		fmt.Println(fmt.Sprintf("Server failed to run"))
		return false
	}
	return true
}

func (httpServer *HTTPServer) StartHTTPServer() bool {
	if len(httpServer.RouteLocationList) == 0 {
		return false
	}
	err := httpServer.HTTPServer.ListenAndServe()
	if err != nil {
		return false
	}
	return true
}
