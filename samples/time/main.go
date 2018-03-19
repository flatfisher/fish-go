package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.ISOWeek())
	fmt.Println(t.Local())
	fmt.Println(t.UTC())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Location())
	fmt.Println(t.Clock())
	fmt.Println(t.Date())
}
