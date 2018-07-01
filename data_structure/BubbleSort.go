package main

import (
	"fmt"
)

func main() {
	slice := []int{12, 23, 536, 564, 23425, 456}
	fmt.Println("UNSORTED LIST ", slice)
	bubblesort(slice)
	fmt.Println("SORETED", slice)

}

func bubblesort(item []int) {
	n := len(item)
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if item[i] > item[i+1] {
				item[i+1], item[i] = item[i], item[i+1]
			}
		}
		if !swapped {
			sorted = true

		}
		n = n - 1
	}
}
