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

func TestAdd(t *testing.T) {
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
	for scenario, fn := range map[string]func(t *testing.T, client http.Client, port int){
		"add 2 numbers": testAdd,
		"sub 2 numbers": testSub,
		"mul 2 numbers": testMul,
		"div 2 numbers": testDiv,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t, client, port)
		})
	}
}

func composeRequest(a float64, b float64) ([]byte, error) {
	reqBody := struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}{
		A: a,
		B: b,
	}
	return json.Marshal(reqBody)
}

func sendRequest(client http.Client, port int, path string, reqBody []byte) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:%v/%v", port, path)
	fmt.Println("url: " + url)
	return client.Post(url, "application/json", bytes.NewBuffer(reqBody))
}

func readResult(resp *http.Response) (float64, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Error httpStatus=%v, response=%v", resp.StatusCode, string(body))
	}
	bodyJson := struct {
		Result float64 `json:"result"`
	}{}
	_ = json.Unmarshal(body, &bodyJson)
	return bodyJson.Result, nil
}

func testAdd(t *testing.T, client http.Client, port int) {
	doTest := func(a float64, b float64, expect float64) {
		req, err := composeRequest(a, b)
		require.NoError(t, err)
		resp, err := sendRequest(client, port, "add", req)
		require.NoError(t, err)
		defer resp.Body.Close()
		rs, err := readResult(resp)
		require.NoError(t, err)
		require.Equal(t, float64(expect), rs)
	}

	doTest(1, 2, 3)
	doTest(2, -1, 1)
}

func testSub(t *testing.T, client http.Client, port int) {
	doTest := func(a float64, b float64, expect float64) {
		req, err := composeRequest(a, b)
		require.NoError(t, err)
		resp, err := sendRequest(client, port, "sub", req)
		require.NoError(t, err)
		defer resp.Body.Close()
		rs, err := readResult(resp)
		require.NoError(t, err)
		require.Equal(t, float64(expect), rs)
	}

	doTest(1, 2, -1)
	doTest(2, -1, 3)
	doTest(100, 1, 99)
}

func testMul(t *testing.T, client http.Client, port int) {
	doTest := func(a float64, b float64, expect float64) {
		req, err := composeRequest(a, b)
		require.NoError(t, err)
		resp, err := sendRequest(client, port, "mul", req)
		require.NoError(t, err)
		defer resp.Body.Close()
		rs, err := readResult(resp)
		require.NoError(t, err)
		require.Equal(t, float64(expect), rs)
	}

	doTest(1, 2, 2)
	doTest(2, -1, -2)
	doTest(100, 1, 100)
	doTest(3, 4, 12)
}

func testDiv(t *testing.T, client http.Client, port int) {
	doTest := func(a float64, b float64, expect float64) {
		req, err := composeRequest(a, b)
		require.NoError(t, err)
		resp, err := sendRequest(client, port, "div", req)
		require.NoError(t, err)
		defer resp.Body.Close()
		rs, err := readResult(resp)
		require.NoError(t, err)
		require.Equal(t, float64(expect), rs)
	}
	doTest(4, 2, 2)
	doTest(0, 2, 0)

	// divide by 0
	req, err := composeRequest(2, 0)
	require.NoError(t, err)
	resp, err := sendRequest(client, port, "div", req)
	require.NoError(t, err)
	defer resp.Body.Close()
	_, err = readResult(resp)
	require.Error(t, err)
}
