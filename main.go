package main

import (
	"fmt"
	"github.com/Liberatys/Sanctuary/service"
)

func main() {
	service := service.NewService("Login", "Login", "A service to handle the login of a user", "3030")
	service.SetDatabaseInformation("127.0.0.1", "3306", "mysql", "root", "Siera_001_DB", "renuo")
	fmt.Println(service.GetDatabaseConnectionString())
	fmt.Println(service.CheckDatabaseConnection())
}
