package networking

type DiscoveryRoutine interface {
	Ping()
	Register()
}

type Discovery struct {
	DiscoveryIP        string
	DiscoveryPort      string
	DiscoverResult     string
	DiscoveryIntervals string
	DiscoveryFailure   bool
}


