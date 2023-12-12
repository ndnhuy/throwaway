package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type LoadBalancer interface {
	// start and return the reverse proxy server that handle incoming requests
	ProxyServer(middleware mux.MiddlewareFunc) (*http.Server, error)
}

func NewLoadBalancer(port int, targetUrls []string) (LoadBalancer, error) {
	services := []ServiceInstance{}
	for _, u := range targetUrls {
		inst, err := NewServiceInstance(u, true)
		if err != nil {
			return nil, err
		}
		services = append(services, inst)
	}
	return &simpleLB{
		services: services,
		port:     port,
	}, nil
}

type simpleLB struct {
	port     int
	services []ServiceInstance
}

func (lb *simpleLB) handleLB(w http.ResponseWriter, r *http.Request) {
	// randomly distribute incoming requests accross services
	i := rand.Uint32() % uint32(len(lb.services))
	lb.services[i].ServeHTTP(w, r)
}

func (lb *simpleLB) ProxyServer(middleware mux.MiddlewareFunc) (*http.Server, error) {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(http.HandlerFunc(lb.handleLB))
	r.Use(middleware)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", lb.port),
		Handler: r,
	}

	return srv, nil
}
