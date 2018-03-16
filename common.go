package main

import (
	"io"
	"net"
	"net/http"
	"encoding/json"
	"time"
)

type Result struct {
	Ret int				`json:"ret"`
	Msg string			`json:"msg"`
	Data []ResultData	`json:"data"`
}

type ResultData struct {
	Url string			`json:"url"`
	StatusCode int		`json:"status"`
	Response string		`json:"response"`
	Err string			`json:"error"`
	Elapsed float64		`json:"elapsed"`
}

func showError(w http.ResponseWriter, ret int, msg string) {
	result := Result{
		Ret :	ret,
		Msg :	msg,
	}
	showResult(w, result)
}

func showSuccess(w http.ResponseWriter, data []ResultData) {
	result := Result{
		Ret :	0,
		Data :	data,
	}
	showResult(w, result)
}

func showResult(w http.ResponseWriter, result Result) {
	b, _ := json.Marshal(result)
	io.WriteString(w, string(b))
}

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns: maxIdleConns,
			MaxIdleConnsPerHost: maxIdleConnsPerHost,
			IdleConnTimeout: time.Duration(idleConnTimeout)* time.Second,
			DisableCompression: true,
		},
		Timeout: 20 * time.Second,
	}
	return client
}
