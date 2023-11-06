package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ndnhuy/mathsys/service"
)

type MiddlewareFunc func(http.Handler) http.Handler

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url:%v%v, duration: %v", r.Host, r.URL.Path, time.Since(startTime).String())
	})
}
func main() {
	srv := NewHTTPServer(":8989")
	log.Fatal(srv.ListenAndServe())
}
func NewHTTPServer(addr string) *http.Server {
	controller := &controller{
		addService: service.AddServiceImpl{},
		subService: service.SubServiceImpl{},
		mulService: service.MultiplyServiceImpl{},
		divService: service.DivServiceImpl{},
	}
	r := mux.NewRouter()
	r.HandleFunc("/add", controller.Add).Methods("POST")
	r.HandleFunc("/sub", controller.Sub).Methods("POST")
	r.HandleFunc("/mul", controller.Mul).Methods("POST")
	r.HandleFunc("/div", controller.Div).Methods("POST")
	r.Use(loggingMiddleware)
	http := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return http
}

type controller struct {
	addService service.AddService
	subService service.SubService
	mulService service.MultiplyService
	divService service.DivService
}

func (c *controller) Add(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// parse request
	var req service.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add
	res := service.AddResponse{Result: c.addService.Do(req.A, req.B)}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%v + %v = %v\n", req.A, req.B, res.Result)
	log.Printf("duration=%v", time.Since(start).String())
}

func (c *controller) Sub(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// parse request
	var req service.SubRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := service.SubResponse{Result: c.subService.Do(req.A, req.B)}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%v - %v = %v\n", req.A, req.B, res.Result)
	log.Printf("duration=%v", time.Since(start).String())
}

func (c *controller) Mul(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// parse request
	var req service.MultiplyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := service.MultiplyResponse{Result: c.mulService.Do(req.A, req.B)}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%v x %v = %v\n", req.A, req.B, res.Result)
	log.Printf("duration=%v", time.Since(start).String())
}

func (c *controller) Div(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// parse request
	var req service.DivRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rs, err := c.divService.Do(req.A, req.B)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := service.DivResponse{Result: rs}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("%v / %v = %v\n", req.A, req.B, res.Result)
	log.Printf("duration=%v", time.Since(start).String())
}
