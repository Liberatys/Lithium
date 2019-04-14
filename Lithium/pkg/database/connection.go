package database

import "database/sql"

type Connection struct {
	Connection *sql.DB
}
