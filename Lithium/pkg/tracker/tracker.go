package tracker

import "time"

type Tracker struct {
	CreationTime           int64
	LastPullTimeStamp      int64
	Requests               int64
	PreviousRequestRecords []int64
}

func InitTracker() Tracker {
	tracker := Tracker{CreationTime: time.Now().Unix(), Requests: 0}
	return tracker
}

func (tracker *Tracker) IncreaseRequests() {
	tracker.Requests++
}

func (tracker *Tracker) GetCurrentRequests() int64 {
	return tracker.Requests
}

func (tracker *Tracker) GetLastRequest() int64 {
	return tracker.PreviousRequestRecords[len(tracker.PreviousRequestRecords)-1]
}

func (tracker *Tracker) StoreCurrentRequests() {
	tracker.PreviousRequestRecords = append(tracker.PreviousRequestRecords, tracker.Requests)
	tracker.Requests = 0
}

func (tracker *Tracker) GetAllRequestRecords() []int64 {
	tracker.LastPullTimeStamp = time.Now().Unix()
	return tracker.PreviousRequestRecords
}
