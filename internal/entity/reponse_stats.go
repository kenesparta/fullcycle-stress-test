package entity

import (
	"sync"
	"time"
)

// ResponseStats holds the count of responses for each status code.
type ResponseStats struct {
	// requests the amount of request that was performed
	requests int

	// statusMap we store the http status code on this map and we count it
	statusMap map[int]int

	// totalDuration the total duration of the request's execution.
	totalDuration time.Duration
	SyMu          sync.Mutex
}

func NewResponseStats() *ResponseStats {
	return &ResponseStats{
		statusMap: make(map[int]int),
	}
}

// IncrementStatusMap increments the count for a given status code.
func (rs *ResponseStats) IncrementStatusMap(statusCode int) {
	rs.SyMu.Lock()
	defer rs.SyMu.Unlock()
	rs.statusMap[statusCode]++
}

func (rs *ResponseStats) CalculateTotalDuration(start, end time.Time) {
	rs.totalDuration = start.Sub(end)
}

func (rs *ResponseStats) IncrementRequest() {
	rs.SyMu.Lock()
	defer rs.SyMu.Unlock()
	rs.requests++
}

func (rs *ResponseStats) TotalDuration() time.Duration {
	return rs.totalDuration
}

func (rs *ResponseStats) Requests() int {
	return rs.requests
}

func (rs *ResponseStats) StatusMap() map[int]int {
	return rs.statusMap
}
