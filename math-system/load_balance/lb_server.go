package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func reqLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// proxyAddr := r.RemoteAddr
		log.Println("req from " + r.RemoteAddr)
		next.ServeHTTP(w, r)
		// targetAddr := r.RemoteAddr
		// log.Printf("%v -> %v", proxyAddr, targetAddr)
	})
}
func main() {
	viper.SetConfigName("services_registry")
	viper.AddConfigPath("/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	serviceUrls := viper.GetStringSlice("services")
	if len(serviceUrls) == 0 {
		log.Fatal("no services found")
	}
	lb, err := NewLoadBalancer(8989, serviceUrls)
	if err != nil {
		log.Fatal(err)
	}
	srv, err := lb.ProxyServer(reqLogger)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
	log.Fatal(srv.ListenAndServe())
}
