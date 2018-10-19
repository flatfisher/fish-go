package main

import "fmt"

func main() {
	var sum float32
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		fmt.Println("good")
	}

	// t := time.Now().UnixNano()
	// rand.Seed(t)
	// s := rand.Intn(6) + 1

	// switch {
	// case s == 1:
	// 	fmt.Println("凶")
	// case s == 2 || s == 3:
	// 	fmt.Println("吉")
	// case s == 4 || s == 5:
	// 	fmt.Println("中吉")
	// case s == 6:
	// 	fmt.Println("大吉")
	// }

	// for i, v := range []int{10, 20, 30} {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }
	// 	var i int
	// LOOP:
	// 	for {
	// 		switch {
	// 		case i == 100:
	// 			break LOOP
	// 		}
	// 		fmt.Println(i)
	// 		i++
	// 	}
}
