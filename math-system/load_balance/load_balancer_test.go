package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
)

func startService(serviceID string) (port int, close func()) {
	ports := dynaport.Get(1)
	port = ports[0]
	host := fmt.Sprintf(":%v", port)
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "I'm "+serviceID)
	})
	srv := &http.Server{
		Addr:    host,
		Handler: r,
	}
	go func() {
		srv.ListenAndServe()
		defer srv.Close()
	}()
	return port, func() { srv.Close() }
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// proxyAddr := r.RemoteAddr
		log.Println("req from " + r.RemoteAddr)
		next.ServeHTTP(w, r)
		// targetAddr := r.RemoteAddr
		// log.Printf("%v -> %v", proxyAddr, targetAddr)
	})
}
func TestLoadBalancer(t *testing.T) {
	port1, close1 := startService("service1")
	defer close1()
	port2, close2 := startService("service2")
	defer close2()
	lbPort := randomPort()
	fmt.Printf("service1:%v\n", port1)
	fmt.Printf("service2:%v\n", port2)
	fmt.Printf("loadbalancer:%v\n", lbPort)
	lb, err := NewLoadBalancer(lbPort, []string{
		fmt.Sprintf("http://localhost:%v", port1),
		fmt.Sprintf("http://localhost:%v", port2),
	})
	require.NoError(t, err)
	srv, err := lb.ProxyServer(loggingMiddleware)
	require.NoError(t, err)
	defer srv.Close()

	go func() {
		srv.ListenAndServe()
	}()

	client := http.DefaultClient
	resp, err := client.Get(fmt.Sprintf("http://localhost:%v/hello", lbPort))
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	fmt.Println(string(body))
}

func randomPort() int {
	ports := dynaport.Get(1)
	port := ports[0]
	return port
}

func TestReadConfig(t *testing.T) {
	viper.SetConfigName("services_registry")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	require.NoError(t, err)

	fmt.Println(viper.Get("services"))
}
