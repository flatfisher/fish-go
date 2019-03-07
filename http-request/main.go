package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://google.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
