package communication

import "net/http"

type Connection interface {
	AddRoute(route string, function func(w http.ResponseWriter, r *http.Request))
	Start()
	SetState(bool)
	GetState() bool
}
