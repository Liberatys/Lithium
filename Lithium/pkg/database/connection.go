// roach wraps `lib/pq` providing the basic methods for
// creating an entrypoint for our database.
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

func New(cfg Config) (roach Connection, err error) {
	roach.cfg = cfg
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslcert=%s sslkey=%s",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port, cfg.ClientCertificate, cfg.ClientKey))
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to postgre database (%s)")
		return
	}

	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err,
			"Couldn't ping postgre database (%s)")
		return
	}
	roach.Db = db
	return
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
