package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "key_from_environment_or_flag")
	q.Add("another_thing", "foo & bar")

	v := url.Values{}
	v.Add("api_key", "key_from_environment_or_flag")
	v.Add("another_thing", "foo & bar")

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
}
