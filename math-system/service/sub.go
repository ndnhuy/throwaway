package service

type SubService interface {
	Do(a float64, b float64) float64
}

type SubServiceImpl struct{}
type SubRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type SubResponse struct {
	Result float64 `json:"result"`
}

func (service SubServiceImpl) Do(a float64, b float64) float64 {
	return a - b
}
