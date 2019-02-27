package main

import (
	"fmt"
	"time"
)

func main() {
	time, _ := time.Parse("2006-01-02T15:04:05+09:00", "2019-02-23T05:16:00+09:00")
	fmt.Println(time)
}
