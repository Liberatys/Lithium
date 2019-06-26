package database

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	Db  *sql.DB
	cfg Config
}

type Config struct {
	Host              string
	Port              string
	User              string
	Password          string
	Database          string
	ClientCertificate string
	ClientKey         string
}

var currentConnection Connection

func InitDatabase(cfg Config) {
	connection, err := New(cfg)
	if err != nil {
		log.Println(err.Error())
	}
	currentConnection = connection
}

func GetDatabaseConnection() Connection {
	return currentConnection
}

func New(cfg Config) (Connection, error) {
	//TODO: find easy way to enable ssl so we can remove ssl disable
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=require",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		return Connection{}, err
	}

	if err = db.Ping(); err != nil {
		return Connection{}, err
	}
	return Connection{db, cfg}, err
}

func (r *Connection) Close() (err error) {
	if r.Db == nil {
		return
	}
	if err = r.Db.Close(); err != nil {
		err = errors.Wrapf(err,
			"Errored closing database connection")
	}
	return
}
