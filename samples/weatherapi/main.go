package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Weather is a weather forecast API response
type Weather struct {
	Title             string `json:"title"`
	Link              string `json:"link"`
	Description       Description
	Location          Location
	PinpointLocations []PinpointLocations
	Forecasts         []Forecasts
	Copyright         Copyright
}

//Description is a weather forecast API response
type Description struct {
	Text       string `json:"text"`
	PublicTime string `json:"publicTime"`
}

//PinpointLocations is a weather forecast API response
type PinpointLocations struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

//Forecasts is a weather forecast API response
type Forecasts struct {
	DateLabel   string `json:"dateLabel"`
	Telop       string `json:"telop"`
	Datte       string `json:"date"`
	Temperature Temperature
	Image       Image
}

//Temperature is a weather forecast API response
type Temperature struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

//Image is a weather forecast API response
type Image struct {
	Width  int    `json:"width"`
	URL    string `json:"url"`
	Title  string `json:"title"`
	Height int    `json:"height"`
}

//Location is a weather forecast API response
type Location struct {
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}

//Copyright is a weather forecast API response
type Copyright struct {
	Provider []Provider
	Link     string `json:"link"`
	Title    string `json:"title"`
	Image    CopyrightImage
}

//Provider is a weather forecast API response
type Provider struct {
	Link string `json:"linl"`
	Name string `json:"name"`
}

//CopyrightImage is a weather forecast API response
type CopyrightImage struct {
	Width  int    `json:"width"`
	Link   string `json:"link"`
	URL    string `json:"url"`
	Title  string `json:"title"`
	Height int    `json:"height"`
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
	fmt.Printf("Results: %v\n", data.Copyright)
	os.Exit(0)
}
