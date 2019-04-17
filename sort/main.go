package main

import (
	"fmt"
)

func bubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	return
}

func selectionSort(a []int) []int {
	for i := range a {
		idx, _ := min(a[i:len(a)])
		a[i], a[i+idx] = a[i+idx], a[i]
	}
	return a
}

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i-j-1] > a[i-j] {
				a[i-j-1], a[i-j] = a[i-j], a[i-j-1]
			} else {
				break
			}
		}
	}
	return a
}

func calcInterval(n int) int {
	h := 1
	for h < n {
		h = 3*h + 1
	}
	h = (h - 1) / 3
	return h
}

func shellSort(a []int) []int {
	h := calcInterval(len(a))
	for h > 1 {
		for i := 0; i < h; i++ {
			b := make([]int, len(a)/h+1)
			cnt := 0
			for j := 0; j < len(a)/h+1; j++ {
				if i+h*j < len(a) {
					b[j] = a[i+h*j]
					cnt++
				}
			}
			c := insertionSort(b[:cnt])
			for j := 0; j < len(c); j++ {
				if i+h*j < len(a) {
					a[i+h*j] = c[j]
				}
			}
		}
		h = (h - 1) / 3
	}
	a = insertionSort(a)
	return a
}

func merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	var i, j, cnt int
	for i+j < len(a)+len(b) {
		if a[i] < b[j] {
			result[cnt] = a[i]
			i++
		} else {
			result[cnt] = b[j]
			j++
		}
		cnt++
		if i == len(a) {
			for j < len(b) {
				result[cnt] = b[j]
				cnt++
				j++
			}
			break
		}
		if j == len(b) {
			for i < len(a) {
				result[cnt] = a[i]
				cnt++
				i++
			}
			break
		}
	}
	return result
}

func divideArray(a []int) ([]int, []int) {
	return a[:len(a)/2], a[len(a)/2:]
}

func mergeSort(a []int) []int {
	a1, a2 := divideArray(a)
	if len(a1) > 1 {
		a1 = mergeSort(a1)
	}
	if len(a2) > 1 {
		a2 = mergeSort(a2)
	}
	a = merge(a1, a2)
	return a
}

func med3(x, y, z int) int {
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		} else {
			return x
		}
	} else {
		if x < z {
			return x
		} else if y < z {
			return z
		} else {
			return y
		}
	}
}

func quickSort(a []int) {
	pivot := med3(a[0], a[len(a)/2], a[len(a)-1])
	left := 0
	right := len(a) - 1
	for {
		for a[left] < pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left >= right {
			break
		}
		a[left], a[right] = a[right], a[left]
		flg := true
		if a[right] == pivot {
			left++
			flg = false
		}
		if a[left] == pivot && flg {
			right--
		}
	}
	a1 := a[:left]
	if len(a1) > 1 {
		quickSort(a1)
	}
	a2 := a[right+1:]
	if len(a2) > 1 {
		quickSort(a2)
	}
	cnt := 0
	for _, v := range a1 {
		a[cnt] = v
		cnt++
	}
	a[cnt] = pivot
	cnt++
	for _, v := range a2 {
		a[cnt] = v
		cnt++
	}
}

func upHeap(a []int, n int) []int {
	a = append(a, n)
	child := len(a) - 1
	var parent = int(child+1)/2 - 1
	for {
		if a[child] > a[parent] {
			a[child], a[parent] = a[parent], a[child]
			child = parent
			parent = (child+1)/2 - 1
			if parent < 0 {
				break
			}
		} else {
			break
		}
	}
	return a
}

func downHeap(a []int) []int {
	a[0], a[len(a)-1] = a[len(a)-1], a[0]
	a = a[:len(a)-1]
	parent := 0
	var child int
	for {
		child = 2*parent + 1
		if child > len(a)-1 {
			break
		}
		if child != len(a)-1 {
			if a[child] < a[child+1] {
				child++
			}
		}
		if a[parent] >= a[child] {
			break
		}
		a[parent], a[child] = a[child], a[parent]
		parent = child
	}
	return a
}

func heapSort(a []int) []int {
	var b []int
	b = append(b, a[0])
	for i := 1; i < len(a); i++ {
		b = upHeap(b, a[i])
	}
	for i := 0; i < len(a); i++ {
		a[len(a)-1-i] = b[0]
		b = downHeap(b)
	}
	return a
}

func main() {
	fmt.Println("Hello World")
}
