package main

import (
	"flag"
	"github.com/kenesparta/fullcycle-stress-test/internal/dto"
)

func setInputFlags() dto.RequestFlagInput {
	var (
		url         = flag.String("url", "", "URL from service to test")
		requests    = flag.Int("r", 10, "Amount of requests to send")
		concurrency = flag.Int("w", 1, "Amount of concurrent requests (workers)")
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
