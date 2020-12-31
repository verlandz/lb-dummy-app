// cmd: go run main.go
// then try to 'curl localhost:8080' / hit 'http://localhost:8080/'
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

type ApiResp struct {
	Count    int64 `json:"count"`
	UnixTime int64 `json:"unix_time"`
}

var (
	// Count how many request that has been through to this app
	COUNT int64 = 0

	// By default APP_PORT is :8080
	// you may change the port by set the ENV
	// ex: APP_PORT=8081 go run main.go
	APP_PORT string = "8080"
)

func init() {
	if os.Getenv("APP_PORT") != "" {
		APP_PORT = os.Getenv("APP_PORT")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	COUNT++

	resp := ApiResp{
		Count:    COUNT,
		UnixTime: time.Now().Unix(),
	}
	b, _ := json.Marshal(resp)

	fmt.Println(string(b))
	fmt.Fprintf(w, string(b))
}

// Don't let /favicon.ico distrub mainHandler
func favHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/favicon.ico", favHandler)
	fmt.Printf("APP_PORT=%v | Golang Ver=%v\n", APP_PORT, runtime.Version())
	fmt.Println(http.ListenAndServe(":"+APP_PORT, nil))
}
