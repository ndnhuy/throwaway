package load_balance

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServiceInstance interface {
	HealthCheck() bool
	getURL() *url.URL
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	// should not expose reverse proxy to outside
	// ReverseProxy() *httputil.ReverseProxy
}

func NewServiceInstance(serviceUrl string, alive bool) (ServiceInstance, error) {
	URL, err := url.Parse(serviceUrl)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(URL)
	return defaultServiceInstance{
		URL:          URL,
		Alive:        alive,
		reverseProxy: proxy,
	}, nil
}

type defaultServiceInstance struct {
	URL          *url.URL
	Alive        bool
	reverseProxy *httputil.ReverseProxy
}

func (s defaultServiceInstance) HealthCheck() bool {
	return true
}

func (s defaultServiceInstance) getURL() *url.URL {
	return s.URL
}

func (s defaultServiceInstance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.reverseProxy.ServeHTTP(w, r)
}

func (s defaultServiceInstance) ReverseProxy() *httputil.ReverseProxy {
	return s.reverseProxy
}
