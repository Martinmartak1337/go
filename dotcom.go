package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://github.com/Martinmartak1337")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
	}
	fmt.Printf("Body: %s", body)
}
