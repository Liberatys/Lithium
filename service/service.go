package service

import (
	"database/sql"
	"fmt"
	"github.com/Liberatys/Sanctuary/communication"
	"github.com/Liberatys/Sanctuary/database"
	"net/http"
)

//Service structure for the base implementation of a service that can be expanded with other "components"
type Service struct {
	Name               string
	Type               string
	Description        string
	IP                 string
	Port               string
	Balancer           ServiceBalancer
	HTTPServer         communication.HTTPConnection
	DatabaseConnection database.DatabaseInformation
}

func NewService(name string, typ string, description string, port string) Service {
	service := Service{
		Name:               name,
		Type:               typ,
		Description:        description,
		Balancer:           ServiceBalancer{},
		Port:               port,
		HTTPServer:         communication.HTTPConnection{},
		DatabaseConnection: database.DatabaseInformation{},
	}
	return service
}

func (service *Service) SetDatabaseInformation(ip string, port string, databasetype string, username string, password string, databasename string) {
	service.DatabaseConnection = database.NewDatabaseInformation(ip, port, databasetype, username, password)
	service.DatabaseConnection.SetDatabaseName(databasename)
	service.DatabaseConnection.Setup()
}

func (service *Service) PrepareQuery(query string, parameters ...string) (*sql.Stmt, error) {
	stmt, stmt_err := service.DatabaseConnection.DatabaseConnection.Prepare(query)
	if stmt_err != nil {
		return nil, stmt_err
	}
	return stmt, nil
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

func (service *Service) SetServiceBalancer(serviceBalancer ServiceBalancer) {
	service.Balancer = serviceBalancer
}

func (service *Service) Start() {
	service.ActivateHTTPServer(service.Port)
}

func (service *Service) ActivateHTTPServer(port string) {
	service.HTTPServer = communication.NewHTTPServer(service.IP, service.Port)
}

func (service *Service) StartHTTPServer() {
	service.HTTPServer.Start()
}

func (service *Service) AddHTTPRoute(route string, fn func(w http.ResponseWriter, r *http.Request)) (bool, string) {
	if service.HTTPServer.Activated {
		service.HTTPServer.AddRoute(route, fn)
		return true, fmt.Sprintf("Route %v added", route)
	}
	return false, "HTTPServer is not activated ... \n Activate using service.ActivateHTTPServer(port)"
}
