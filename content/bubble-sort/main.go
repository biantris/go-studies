package main

import "fmt"

func bubbleSort(n []int) []int {
	for i, _ := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	return n
}

func main() {
	unsorted := []int{1, 5, 4, 3, 2}
	fmt.Println(bubbleSort(unsorted))
}