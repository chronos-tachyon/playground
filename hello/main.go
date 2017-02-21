package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var flagListen = flag.String("listen", ":8000", "host:port to listen on")
var hostname string

type answer struct {
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
}

type hello struct{}

func (_ *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

type healthz struct{}

func (_ *healthz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Length", "0")
	w.WriteHeader(200)
}

func main() {
	flag.Parse()

	cmd := exec.Command("hostname", "--fqdn")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	hostname = strings.TrimSpace(string(out))

	log.Printf("Hostname: %s", hostname)
	log.Printf("Listening on: %s", *flagListen)

	http.Handle("/", &hello{})
	http.Handle("/healthz", &healthz{})
	http.ListenAndServe(*flagListen, nil)
}
