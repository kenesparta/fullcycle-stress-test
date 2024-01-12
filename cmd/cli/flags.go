package main

import (
	"flag"
	"github.com/kenesparta/fullcycle-stress-test/internal/dto"
)

func setInputFlags() dto.RequestFlagInput {
	var (
		url         = flag.String("url", "", "URL from service to test")
		requests    = flag.Int("r", 1, "Amount of requests to send")
		concurrency = flag.Int("c", 1, "Time in seconds of each request")
		method      = flag.String("m", "GET", "HTTP method to use")
	)

	flag.Parse()
	reqFl := dto.RequestFlagInput{
		RequestAmount: *requests,
		Concurrency:   *concurrency,
		URL:           *url,
		Method:        *method,
	}

	return reqFl
}
