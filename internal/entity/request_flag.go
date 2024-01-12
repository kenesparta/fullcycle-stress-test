package entity

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RequestFlag struct {
	RequestAmount int
	Concurrency   int
	URL           string
	Method        string
}

func (rf *RequestFlag) Validate() error {
	if rf.RequestAmount <= 0 {
		return fmt.Errorf("request should be positive number")
	}

	if rf.Concurrency <= 0 {
		return fmt.Errorf("concurrency value should be positive number")
	}

	if rf.Concurrency > rf.RequestAmount {
		return fmt.Errorf("we have more workers (go routines) than requests")
	}

	if !isValidURL(rf.URL) {
		return fmt.Errorf("should be a URL valid value")
	}

	if !isMethodAllowed(rf.Method) {
		return fmt.Errorf("should be a valid HTTP method")
	}

	return nil
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}

	return (parsedURL.Scheme == "http" ||
		parsedURL.Scheme == "https") &&
		parsedURL.Host != ""
}

func isMethodAllowed(m string) bool {
	mUp := strings.ToUpper(m)
	return mUp == http.MethodGet ||
		mUp == http.MethodPost ||
		mUp == http.MethodOptions
}
