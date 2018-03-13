package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Print("started.")

	funcs := []func(){
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
		},
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
		},
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
		},
	}

	var waitGroup sync.WaitGroup

	for _, sleep := range funcs {
		waitGroup.Add(1)

		go func(function func()) {
			defer waitGroup.Done()
			function()
		}(sleep)
	}

	waitGroup.Wait()

	log.Print("all finished.")
}
