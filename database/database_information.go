package database

import "database/sql"

//TODO: imporve the way the connection to the database is handled

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
		return "You first have to set the name of the database"
	}
	CreateDatabaseConnection(info)
	info.DatabaseConnection = GetDatabaseConnection()
	return "Database connection established"
}

func (info *DatabaseInformation) Close() error {
	err := info.DatabaseConnection.Close()
	return err
}

func (info *DatabaseInformation) SetDatabaseName(name string) {
	info.DatabaseName = name
}
