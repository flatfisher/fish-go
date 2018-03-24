package main

import "fmt"

func getHello() string {
	return "Hello World"
}

func getWorld() string {
	return "World"
}

func main() {
	fmt.Println(getHello())
}
