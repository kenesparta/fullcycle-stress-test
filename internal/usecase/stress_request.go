package usecase

type StressRequest struct {
}

func NewStressRequest() *StressRequest {
	return &StressRequest{}
}

func (sr *StressRequest) Execute() {}
