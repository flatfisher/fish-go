package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000 / 1.0
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(int64(n))
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(math.Sin(n))
}
