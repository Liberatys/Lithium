package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPServer struct {
	RouteLocationList []string
	Router            *mux.Router
	HTTPServer        *http.Server
	Port              string
}

//TODO: Create method of removing routes

func InitializeBaiscHTTTPServer(Port string) HTTPServer {
	httpServer := HTTPServer{Router: mux.NewRouter().StrictSlash(true),Port:Port}
	//setting a low repsonse time for server and client, should solve the problem of DDOS with blocking like slow loris.
	timeOut := time.Second * 1
	server := &http.Server{
		Addr:         ":" + Port,
		ReadTimeout:  timeOut,
		WriteTimeout: timeOut,
		IdleTimeout:  timeOut,
		Handler:      &CORSRouterDecorator{httpServer.Router},
	}
	httpServer.HTTPServer = server
	return httpServer
}

func (httpServer *HTTPServer) AddRoute(RouteLocation string, RouteFunction func(w http.ResponseWriter, req *http.Request)) {
	httpServer.RouteLocationList = append(httpServer.RouteLocationList, RouteLocation)
	httpServer.Router.HandleFunc(RouteLocation, RouteFunction)
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
