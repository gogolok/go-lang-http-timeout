package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

func LongSleepHandler(w http.ResponseWriter, r *http.Request) {
	for {
	}
	io.WriteString(w, "Long sleeping")
}

func webserver() {
	// start web server that sleeps for a long time per request
	http.HandleFunc("/long-sleep", LongSleepHandler)
	http.ListenAndServe(":3001", nil)
}

func requests() {
	fmt.Println("Start: ", time.Now())

	request := gorequest.New()
	resp, _, errs := request.Get("http://localhost:3001/long-sleep").End()
	if errs != nil {
		fmt.Println("errs[0] = %v", errs[0])
		return
	}
	fmt.Println("Response ", resp)
	fmt.Println("Stop: ", time.Now())
	fmt.Println("Request finished.")
}

func main() {
	fmt.Println("Start webserver.")
	go webserver()

	time.Sleep(1 * time.Second)

	go requests()

	for {
	}
}
