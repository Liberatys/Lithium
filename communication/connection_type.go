package communication

//Type is an interface for the different connection options that can be used with a service
type ConnectionType interface {
	Start()
}
