package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var connection Connection

type Connection struct {
	DatabaseType       string
	ConnectionSequence string
	Connection         *sql.DB
}

func InitConnection(databaseType string, username string, password string, databaseName string, protocol string, port string, IP string) {
	connectionSequence := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, protocol, IP, port, databaseName)
	connection := Connection{DatabaseType: databaseType, ConnectionSequence: connectionSequence}
	createNewConnection(&connection)
}

func createNewConnection(connection *Connection) *sql.DB {
	var err error
	connection.Connection, err = sql.Open(connection.DatabaseType, connection.ConnectionSequence)
	if err != nil {
		log.Fatal("Creation of Database connection failed")
		return nil
	} else {
		return connection.Connection
	}
}

func GetDatabaseConnection() *sql.DB {
	if connection.Connection != nil && IsConnected() == true {
		return connection.Connection
	}
	if connection.Connection == nil {
		connection.Connection = createNewConnection(&connection)
	}
	return connection.Connection
}

func IsConnected() bool {
	err := connection.Connection.Ping()
	if err != nil {
		return false
	}
	return true
}

func RetreaveConnectionInformation() Connection {
	return connection
}
