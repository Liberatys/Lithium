package State

import (
	"database/sql"
	"time"
)

type ServiceState struct {
	CreationTimeStamp  int64
	DatabaseType       string
	DatabaseConnection *sql.DB
	Status             string
}

func (serviceState *ServiceState) InitServiceState(DatabaseType string) {
	serviceState.CreationTimeStamp = time.Now().Unix()
}

func (serviceState *ServiceState) CheckStatus() string {
	serviceState.Status = "Working Fine"
	err := serviceState.DatabaseConnection.Ping()
	if err != nil {
		serviceState.Status = "Database not working"
	}
	return serviceState.Status
}

func (serviceState *ServiceState) PingDatabase() {
	serviceState.DatabaseConnection.Ping()
}
