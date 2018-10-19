package main

import "fmt"

func calculateTriangleArea(base float32, height float32) float32 {
	area := (base * height) / 2
	return area
}

func main() {
	fmt.Print(calculateTriangleArea(10, 20))
}
