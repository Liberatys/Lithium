package database

import "database/sql"

type DatabaseInformation struct {
	IP                 string
	Port               string
	DatabaseTyp        string
	Username           string
	Password           string
	DatabaseName       string
	DatabaseConnection *sql.DB
}

func NewDatabaseInformation(ip string, port string, databasetype string, username string, password string) DatabaseInformation {
	information := DatabaseInformation{
		IP:          ip,
		Port:        port,
		DatabaseTyp: databasetype,
		Username:    username,
		Password:    password,
	}
	return information
}

func (info *DatabaseInformation) Setup() string {
	if info.DatabaseName == "" {
		return "Please set the databasename first"
	}
	CreateDatabaseConnection(info)
	info.DatabaseConnection = GetDatabaseConnection()
	return "Databaseconnection established"
}

func (info *DatabaseInformation) Close() error {
	err := info.DatabaseConnection.Close()
	return err
}

func (info *DatabaseInformation) SetDatabaseName(name string) {
	info.DatabaseName = name
}
