package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

func printWorld() {
	fmt.Println("World")
}

func sendHello(ch chan<- string) {
	ch <- "Hello"
}

func main() {
	// printHello()
	// printWorld()
	// time.Sleep(time.Second)

	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go printHello(&wg)
	// }
	// wg.Wait()

	ch := make(chan string)
	go sendHello(ch)

	msg := <-ch
	fmt.Println(msg)

	q := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		q <- "完了"
		close(q)
	}()

	// n := 0
	// loop:
	// 	for {
	// 		select {
	// 		case msg := <-q:
	// 			fmt.Println(msg)
	// 			break loop
	// 		default:
	// 		}
	// 		n++
	// 		fmt.Println("何か", n)
	// 		time.Sleep(1 * time.Second)
	// 	}

	n := 0
	for {
		select {
		case msg := <-q:
			fmt.Println(msg)
		case <-time.After(3 * time.Second):
			fmt.Println("タイムアウト")
		default:
		}
		n++
		fmt.Println("何か", n)
		time.Sleep(1 * time.Second)
	}
}
