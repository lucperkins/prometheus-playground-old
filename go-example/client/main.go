package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	hostAndPort = "web:2112"
)

var (
	host string
	port int
)

func runQuery(addr string) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Fatalf("Could not create HTTP request: %v", err)
	}

	q := req.URL.Query()
	q.Add("name", "Luc")
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatalf("Request error: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not parse response: %v", err)
	}

	fmt.Println(string(body))
}

func main() {
	host := os.Getenv("WEB_HOST")
	port := os.Getenv("WEB_PORT")

	addr := fmt.Sprintf("http://%s:%s/greeting", host, port)

	i := 0

	for i < 1000 {
		runQuery(addr)
		i++
	}

	os.Exit(0)
}
