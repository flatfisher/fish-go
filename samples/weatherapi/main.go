package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Weather struct {
	Title string `json:"title"`
}

func main() {
	url := "http://weather.livedoor.com/forecast/webservice/json/v1?city=040010"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data Weather
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)
	os.Exit(0)
}
