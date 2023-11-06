package service

type MultiplyService interface {
	Do(a float64, b float64) float64
}

type MultiplyServiceImpl struct{}
type MultiplyRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type MultiplyResponse struct {
	Result float64 `json:"result"`
}

func (service MultiplyServiceImpl) Do(a float64, b float64) float64 {
	return a * b
}
