package infra

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/kenesparta/fullcycle-stress-test/internal/entity"
)

type MakeHttpRequest struct{}

func NewMakeHttpRequest() *MakeHttpRequest {
	return &MakeHttpRequest{}
}

func (mh *MakeHttpRequest) Get(rf entity.RequestFlag, cm entity.ConcurrencyMgmt) {
	req, reqErr := http.NewRequest(rf.Method, rf.URL, nil)
	if reqErr != nil {
		cm.Err <- reqErr
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	response, doErr := client.Do(req)
	if doErr != nil {
		if response != nil {
			cm.Status <- response.StatusCode
			return
		}

		cm.Err <- doErr
		return
	}

	defer func(Body io.ReadCloser) {
		if bcErr := Body.Close(); bcErr != nil {
			cm.Err <- bcErr
			return
		}
		<-cm.ReleaseSpot
		cm.Done <- true
	}(response.Body)

	cm.Status <- response.StatusCode
}
