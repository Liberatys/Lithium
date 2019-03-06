package Service

import (
	"github.com/Liberatys/Lithium/Networking"
	"github.com/Liberatys/Lithium/State"
	"log"
	"net/http"
	"time"
)

type ServiceInterface interface {
	GetName() string
	GetType() string
	GetDescription() string
	Ping() float64
	Alive() bool
	Response() http.Response
	Request(r http.Request)
	Initialize()
	GetRoutes()
}

type Service struct {
	Name        string
	Type        string
	Description string
	IP          string
	Port        string
	PitchPoint  bool
	Status      State.ServiceState
	Router      ServiceRouter
	HomePoint   HomePoint
	ErrorCount  int
}

type HomePoint struct {
	HomeIP       string
	HomePort     string
	HomeEndPoint string
}

func NewService(Name string, Type string, Description string, Port string, PitchPoint bool) Service {
	newService := Service{Name: Name, Type: Type, Description: Description, Port: Port, IP: Networking.RetreaveCurrentOutpoundIP().String(), PitchPoint: PitchPoint}
	return newService
}

/*
	Creates the information pool for the homeserver adress and port.
	Creates the other systems that are needed for the service to run.
	Afther that, it opens the connections, and calls to the server.
*/
func (service *Service) Initialize(HomeIP string, HomePort string, HomeEndPoint string) {
	log.Println("Service " + service.Name + " created and initialized")
	service.HomePoint = HomePoint{HomeIP: HomeIP, HomePort: HomePort, HomeEndPoint: HomeEndPoint}
	service.setupServiceRouter()
	service.setupServiceState("mysql")
}

func (service *Service) AddRoute(endPointPath string, function func(http.ResponseWriter, *http.Request)) {
	service.Router.registerRoute(endPointPath, function)
}

func (service *Service) setupServiceState(DatabaseType string) {
	service.Status = State.ServiceState{}
	service.Status.InitServiceState(DatabaseType)
}

func (service *Service) setupServiceRouter() {
	service.Router = ServiceRouter{}
	service.Router.InitServiceRouter()
}

func (service *Service) OpenForConnection() {
	log.Println("Opening connection for service")
	if service.PitchPoint == true {
		service.PitchServer()
	}
	service.Router.OpenConnections()
	timeOutInSeconds := time.Second * 1
	server := &http.Server{
		Addr:         ":" + service.Port,
		ReadTimeout:  timeOutInSeconds,
		WriteTimeout: timeOutInSeconds,
		IdleTimeout:  timeOutInSeconds,
		Handler:      &CORSRouterDecorator{&service.Router.Router},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Service start failed....")
	}
}

func (service *Service) GetName() string {
	return service.Name
}

func (service *Service) GetType() string {
	return service.Type
}

func (service *Service) GetDescription() string {
	return service.Description
}
