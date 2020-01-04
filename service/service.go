package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Liberatys/Sanctuary/communication"
	"github.com/Liberatys/Sanctuary/database"
	"github.com/Liberatys/Sanctuary/load"
)

//Service structure for the base implementation of a service that can be expanded with other "components"
type Service struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Description        string `json:"description"`
	IP                 string `json:"ip"`
	Port               string `json:"port"`
	DefaultRoutes      bool
	Balancer           ServiceBalancer
	HTTPServer         communication.Connection
	DatabaseConnection database.DatabaseInformation
}

func (service *Service) Serialize() string {
	serviceSer := communication.SerializedService{
		Name:        service.Name,
		Type:        service.Type,
		Description: service.Description,
		IP:          service.IP,
		Port:        service.Port,
		Load:        load.NewLoad(),
	}
	return communication.Serialize(serviceSer)
}

func NewService(name string, typ string, description string, port string) Service {
	service := Service{
		Name:               name,
		Type:               typ,
		Description:        description,
		Balancer:           ServiceBalancer{},
		Port:               port,
		HTTPServer:         &communication.HTTPConnection{},
		DatabaseConnection: database.DatabaseInformation{},
		DefaultRoutes:      true,
	}
	return service
}

func (service *Service) DisableDefaultRoutes() {
	service.DefaultRoutes = false
}

func (service *Service) SetDatabaseInformation(ip string, port string, databasetype string, username string, password string, databasename string) string {
	service.DatabaseConnection = database.NewDatabaseInformation(ip, port, databasetype, username, password)
	service.DatabaseConnection.SetDatabaseName(databasename)
	return service.DatabaseConnection.Setup()
}

func (service *Service) GetDatabaseConnection() *sql.DB {
	return service.DatabaseConnection.DatabaseConnection
}

func (service *Service) ExecutePerparedQuery(query string, parameters ...interface{}) (sql.Result, error) {
	stmt, stmt_err := service.DatabaseConnection.DatabaseConnection.Prepare(query)
	if stmt_err != nil {
		return nil, stmt_err
	}
	return stmt.Exec(parameters...)
}

func (service *Service) SerializeQueryResult(statement sql.Result) string {
	return ToJSON(statement)
}

func ToJSON(obj interface{}) string {
	res, err := json.Marshal(obj)
	if err != nil {
		panic("error with json serialization " + err.Error())
	}
	return string(res)
}

func (service *Service) Shutdown() error {
	if service.DatabaseConnection.DatabaseName != "" {
		return service.DatabaseConnection.Close()
	}
	return fmt.Errorf("no errors accured")
}

func (service *Service) CheckDatabaseConnection() bool {
	if service.DatabaseConnection.DatabaseName == "" {
		return false
	}
	err := service.DatabaseConnection.DatabaseConnection.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (service *Service) GetDatabaseConnectionString() string {
	if service.DatabaseConnection.DatabaseName == "" {
		return "Not Connected"
	}
	databaseInfo := service.DatabaseConnection
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", databaseInfo.Username, databaseInfo.Password, databaseInfo.IP, databaseInfo.Port, databaseInfo.DatabaseName)
}

func (service *Service) SetServiceBalancer(ip, port string) {
	service.Balancer = NewServiceBalancer(ip, port, "")
}

func (service *Service) Register() {
	if service.Balancer.Configured == true {
		err, message := service.Balancer.Connect(service)
		if err != nil {
			panic(message)
		}
	}
}

func (service *Service) DefaultStartUp() {
	if service.DefaultRoutes == false {
		panic("Default Routes must be enabled for DefaultStartup")
	}
	service.ActivateHTTPServer()
	service.AddDefaultRoutes()
	service.Register()
	service.StartHTTPServer()
}

func (service *Service) AddDefaultRoutes() {
	service.HTTPServer.AddRoute("/healthcheck", HealthCheck)
	service.HTTPServer.AddRoute("/notification", Notification)
}

func (service *Service) ActivateHTTPServer() {
	service.HTTPServer = communication.NewHTTPServer(service.IP, service.Port)
	if service.DefaultRoutes == true {
		service.AddDefaultRoutes()
	}
}

func (service *Service) ActivateHTTPSServer(cert, key string) {
	server := communication.NewHTTPSServer(service.IP, service.Port)
	server.Cert = cert
	server.Key = key
	service.HTTPServer = server
	if service.DefaultRoutes == true {
		service.AddDefaultRoutes()
	}
}

func (service *Service) StartHTTPServer() {
	service.HTTPServer.Start()
}

func (service *Service) AddHTTPRoute(route string, fn func(w http.ResponseWriter, r *http.Request)) (bool, string) {
	if service.HTTPServer.GetState() {
		service.HTTPServer.AddRoute(route, fn)
		return true, fmt.Sprintf("Route %v added", route)
	}
	return false, "HTTPServer is not activated ... \n Activate using service.ActivateHTTPServer(port)"
}
