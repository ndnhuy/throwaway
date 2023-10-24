package service

type AddService interface {
	Do(a float64, b float64) float64
}

type AddServiceImplm struct{}
type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}

func (service AddServiceImplm) Do(a float64, b float64) float64 {
	return a + b
}
