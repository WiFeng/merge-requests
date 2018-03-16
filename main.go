package main

import (
	"flag"
	"log"
	"net/http"
//	"runtime"
//	_ "net/http/pprof"
)

var (
	addr string
	maxIdleConns int
	maxIdleConnsPerHost int
	idleConnTimeout int

	httpClient *http.Client
)

func init() {

	flag.StringVar(&addr, "addr", ":8080", "http service address")
	flag.IntVar(&maxIdleConns, "maxIdleConns", 500, "http MaxIdleConns")
	flag.IntVar(&maxIdleConnsPerHost, "maxIdleConnsPerHost", 150, "http MaxIdleConnsPerHost")
	flag.IntVar(&idleConnTimeout, "idleConnTimeout", 90, "http IdleConnTimeout")

	flag.Parse()
	
	// runtime.GOMAXPROCS(20);

	// init HTTPClient
	httpClient = createHTTPClient()
}

func main() {
	// api handler
	http.HandleFunc("/merge", serveMerge)
	// listen
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
