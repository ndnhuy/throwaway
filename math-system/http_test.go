package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/travisjeffery/go-dynaport"
)

func TestFoo(t *testing.T) {
	ports := dynaport.Get(1)
	port := ports[0]
	host := fmt.Sprintf(":%v", port)
	server := NewHTTPServer(host)

	go func() {
		err := server.ListenAndServe()
		require.NoError(t, err)
		defer server.Close()
	}()

	client := http.Client{}
	reqBody := struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}{
		A: 1,
		B: 2,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	resp, err := client.Post(fmt.Sprintf("http://localhost:%v/", port), "application/json", bytes.NewBuffer(reqBodyJson))
	require.NoError(t, err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	bodyJson := struct {
		Result float64 `json:"result"`
	}{}
	_ = json.Unmarshal(body, &bodyJson)
	require.Equal(t, float64(3), bodyJson.Result)
}
