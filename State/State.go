package State

import "database/sql"

type ServiceState struct {
	CreationTimeStamp  int64
	DatabaseConnection *sql.DB
	Status             string
}

func (serviceState *ServiceState) checkStatus() {
	serviceState.Status = "Working Fine"
	err := serviceState.DatabaseConnection.Ping()
	if err != nil {
		serviceState.Status = "Database not working"
	}
}
