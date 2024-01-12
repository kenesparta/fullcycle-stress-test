package dto

type RequestFlagInput struct {
	RequestAmount int
	Concurrency   int
	URL           string
	Method        string
}

type RequestFlagOutput struct {
	Requests      int         `json:"requests"`
	Workers       int         `json:"workers"`
	StatusMap     map[int]int `json:"status_map_count"`
	Errors        []string    `json:"errors"`
	ErrorCount    int         `json:"error_count"`
	TotalDuration string      `json:"total_duration"`
}
