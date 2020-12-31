// cmd: go run main.go
// then try to 'curl localhost:8080' / hit 'http://localhost:8080/'
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

type ApiResp struct {
	Count    int64  `json:"count"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	UnixTime int64  `json:"unix_time"`
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

// GetLocalIP returns the non loopback local IP of the host
// source: https://stackoverflow.com/a/31551220
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	COUNT++

	HOSTNAME, _ := os.Hostname()
	resp := ApiResp{
		Count:    COUNT,
		Hostname: HOSTNAME,
		IP:       GetLocalIP(),
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
