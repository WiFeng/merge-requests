package main

import (
	"io/ioutil"
//	"log"
//	"encoding/json"
	"net/http"
	"time"

//	 "fmt"
)

type ChanMsg struct {
	Id int
	Resp *http.Response
	Err error
	Elapsed time.Duration
}

func serveMerge(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	urls := query["urls"]
	length := len(urls)

	if length < 1 {
		showError(w, 10001, "urls is required")
		return
	
	}
	
	c := make(chan ChanMsg)
	d := make([]ResultData, length)

	for id, url := range urls {
		go func(id int, url string) {
			start := time.Now()
			resp, err := httpClient.Get(url)
			elapsed := time.Since(start)
			// fmt.Println("elapsed:", elapsed)
			c <- ChanMsg{id, resp, err, elapsed}
		}(id, url)
	}


	for len := length; len > 0; len-- {
		
		var id int
		var resp *http.Response
		var err error
		var data ResultData
		var elapsed time.Duration
		
		select {
		case comResp, ok := <- c :
			if !ok {
				showError(w, 10002, "error")
				return
			}

			id = comResp.Id
			resp = comResp.Resp
			err = comResp.Err
			elapsed = comResp.Elapsed
			if err != nil {
				data = ResultData{
					Url : urls[id],
					Err : err.Error(),
					Elapsed : elapsed.Seconds(),
				}
			} else {
				statusCode := resp.StatusCode
				body, _ := ioutil.ReadAll(resp.Body)

				data = ResultData{
					Url : urls[id],
					StatusCode : statusCode,
					Response : string(body),
					Elapsed : elapsed.Seconds(),
				}
				
				resp.Body.Close()
			}
			// fmt.Println(id, urls[id])
			d[id] = data
		}

	}

	showSuccess(w, d);
}
