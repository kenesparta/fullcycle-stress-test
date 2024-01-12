package entity

// ConcurrencyMgmt is a struct that facilitates managing concurrent operations in Go.
type ConcurrencyMgmt struct {
	// Status is a channel for sending status codes (int) from each goroutine and per HTTP Request.
	Status chan int

	// Done is a channel used to signal the completion of a goroutine's task.
	Done chan bool

	// Err it stores the error channel
	Err chan error

	// ReleaseSpot is a channel that acts as a semaphore for controlling concurrency.
	ReleaseSpot chan struct{}
}
