package service

import "errors"

type DivService interface {
	Do(a float64, b float64) (float64, error)
}

type DivServiceImpl struct{}
type DivRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type DivResponse struct {
	Result float64 `json:"result"`
}

func (service DivServiceImpl) Do(a float64, b float64) (float64, error) {
	if b == 0 {
		return -1, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
