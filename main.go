package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
)

type hello struct{}

type answer struct {
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
}

func (_ *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		names, err := net.LookupAddr(addr)
		if err == nil {
			hostname = names[0]
			break
		}
	}

	a := answer{
		Hostname: hostname,
		Message:  "Hello, world!",
	}
	data, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	data = append(data, '\r', '\n')

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/", &hello{})
	http.ListenAndServe(":8000", nil)
}
