package load

type Load struct {
	Network int64
	CPU     int64
}

func NewLoad() Load {
	load := getNewLoadReading()
	return load
}
func getNewLoadReading() Load {
	// loading the detail on the network load as well as the cpu load
	return Load{}
}
