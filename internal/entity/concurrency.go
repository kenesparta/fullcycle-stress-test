package entity

type Concurrency struct {
	Status chan int
	Done   chan bool
	Spot   chan struct{}
}
