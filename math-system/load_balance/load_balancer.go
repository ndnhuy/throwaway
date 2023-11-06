package load_balance

import (
	"fmt"
	"math/rand"
	"net/http"
)

type LoadBalancer interface {
	// start and return the reverse proxy server that handle incoming requests
	StartProxyServer() (*http.Server, error)
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

func (lb *simpleLB) StartProxyServer() (*http.Server, error) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", lb.port),
		Handler: http.HandlerFunc(lb.handleLB),
	}
	go func() {
		srv.ListenAndServe()
		srv.Close()
	}()

	return srv, nil
}
