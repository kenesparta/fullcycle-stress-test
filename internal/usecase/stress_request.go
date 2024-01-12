package usecase

import (
	"fmt"
	"time"

	"github.com/kenesparta/fullcycle-stress-test/internal/dto"
	"github.com/kenesparta/fullcycle-stress-test/internal/entity"
)

type StressRequest struct {
	statsRepo entity.StatsRepo
	presenter Presenter
}

func NewStressRequest(statsRepo entity.StatsRepo, presenter Presenter) *StressRequest {
	return &StressRequest{
		statsRepo: statsRepo,
		presenter: presenter,
	}
}

func (sr *StressRequest) Execute(flag dto.RequestFlagInput) ([]byte, error) {
	rf := entity.RequestFlag{
		RequestAmount: flag.RequestAmount,
		Concurrency:   flag.Concurrency,
		URL:           flag.URL,
		Method:        flag.Method,
	}
	if valErr := rf.Validate(); valErr != nil {
		return nil, valErr
	}

	var (
		statusChan      = make(chan int, rf.RequestAmount)
		doneChan        = make(chan bool, rf.RequestAmount)
		errChan         = make(chan error, rf.RequestAmount)
		concurrencyChan = make(chan struct{}, rf.Concurrency)
	)

	startTime := time.Now()
	for i := 0; i < rf.RequestAmount; i++ {
		go sr.statsRepo.Get(
			rf,
			entity.ConcurrencyMgmt{
				Status:      statusChan,
				Done:        doneChan,
				ReleaseSpot: concurrencyChan,
				Err:         errChan,
			},
		)
	}

	respStats := entity.NewResponseStats()
	for i := 0; i < rf.RequestAmount; i++ {
		select {
		case status := <-statusChan:
			respStats.IncrementRequest()
			respStats.IncrementStatusMap(status)
		case errRes := <-errChan:
			respStats.IncrementRequest()
			respStats.AddingErrors(errRes)
		case <-doneChan:
		}
	}

	respStats.CalculateTotalDuration(startTime, time.Now())
	reqErrors := func() []string {
		errMsgList := make([]string, 0)
		for _, e := range respStats.Errors() {
			errMsgList = append(errMsgList, e.Error())
		}
		return errMsgList
	}()

	return sr.presenter.Present(
		dto.RequestFlagOutput{
			Requests:      respStats.Requests(),
			Workers:       rf.Concurrency,
			StatusMap:     respStats.StatusMap(),
			TotalDuration: fmt.Sprintf("%s", respStats.TotalDuration()),
			Errors:        reqErrors,
			ErrorCount:    len(reqErrors),
		},
	)
}
