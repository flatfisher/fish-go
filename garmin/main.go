package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Garmin manage weiht
type Garmin struct {
	Date   int64   `json:"date"`
	Weight float64 `json:"weight"`
}

func main() {
	f, err := ioutil.ReadFile("./garmin.json")
	if err != nil {
		os.Exit(1)
	}

	var garmins []Garmin

	json.Unmarshal(f, &garmins)

	for _, garmin := range garmins {
		fmt.Println(time.Unix(garmin.Date/1000, 0).Format("2006/01/02 15:04:05"))
		fmt.Printf("%.1f kg\n", garmin.Weight/1000)
	}
}
