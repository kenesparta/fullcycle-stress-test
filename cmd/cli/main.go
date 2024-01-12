package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/kenesparta/fullcycle-ratelimiter/internal/entity"
)

func main() {
	var (
		url         = flag.String("url", "", "URL from service to test")
		requests    = flag.Int("r", 1, "Amount of requests to send")
		concurrency = flag.Int("c", 1, "Time in seconds of each request")
		method      = flag.String("m", "GET", "HTTP method to use")
	)
	flag.Parse()

	rf := entity.RequestFlag{
		RequestAmount: *requests,
		Concurrency:   *concurrency,
		URL:           *url,
		Method:        *method,
	}

	if valErr := rf.Validate(); valErr != nil {
		fmt.Println(valErr.Error())
		return
	}

	statusChan := make(chan int, rf.RequestAmount)
	doneChan := make(chan bool, rf.RequestAmount)
	concurrencyChan := make(chan struct{}, rf.Concurrency)

	startTime := time.Now()

	for i := 0; i < rf.RequestAmount; i++ {
		go sendRequest(rf.URL, rf.Method, statusChan, doneChan, concurrencyChan)
	}

	statusCounts := make(map[int]int)
	for i := 0; i < rf.RequestAmount; i++ {
		select {
		case status := <-statusChan:
			statusCounts[status]++
		case <-doneChan:
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Println("Completed all requests. Response status codes:")
	for status, count := range statusCounts {
		if status == 0 {
			fmt.Printf("Failed requests: %d\n", count)
		} else {
			fmt.Printf("Status Code %d: %d requests\n", status, count)
		}
	}

	fmt.Printf("Duration of the test: %s\n", duration)
}

func sendRequest(
	url string,
	method string,
	statusChan chan<- int,
	doneChan chan<- bool,
	concurrencySpot chan struct{},
) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Error creating request: %s\n", err)
		statusChan <- 0
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Request error: %s\n", err)
		statusChan <- 0
		return
	}

	defer func(Body io.ReadCloser) {
		if bcErr := Body.Close(); bcErr != nil {
			log.Printf("%s\n", bcErr.Error())
			return
		}
		<-concurrencySpot
		doneChan <- true
	}(response.Body)

	statusChan <- response.StatusCode
}
