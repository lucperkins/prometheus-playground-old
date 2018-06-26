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

func main() {
	resp, err := http.Get(fmt.Sprintf("http://%s", hostAndPort))

	if err != nil {
		log.Fatalf("Could not make GET request to service: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Could not parse response: %v", err)
	}

	fmt.Println(string(body))

	os.Exit(0)
}
