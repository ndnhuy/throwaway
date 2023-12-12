package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
)

func TestReverseProxy(t *testing.T) {
	// start math server
	tearDown, port := initServer()
	defer tearDown()

	targetURL, err := url.Parse(fmt.Sprintf("http://localhost:%v", port))
	require.NoError(t, err)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxyPort := randomPort()
	go func() {
		proxyServer := &http.Server{
			Addr:    fmt.Sprintf(":%v", proxyPort),
			Handler: proxy,
		}
		_ = proxyServer.ListenAndServe()
		defer proxyServer.Close()
	}()

	client := http.Client{}
	testAdd(t, client, proxyPort)
}

func randomPort() int {
	ports := dynaport.Get(1)
	port := ports[0]
	return port
}
