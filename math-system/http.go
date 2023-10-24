package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndnhuy/mathsys/service"
)

func main() {
	srv := NewHTTPServer(":8989")
	log.Fatal(srv.ListenAndServe())
}
func NewHTTPServer(addr string) *http.Server {
	controller := &controller{
		addService: service.AddServiceImplm{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Add).Methods("POST")
	http := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return http
}

type controller struct {
	addService service.AddService
}

type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}

func (c *controller) Add(w http.ResponseWriter, r *http.Request) {
	// parse request
	var req AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add
	res := AddResponse{Result: c.addService.Do(req.A, req.B)}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
