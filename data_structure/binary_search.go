// Binary search in go
package main

import (
	"fmt"
	"sort"
)

func BinarySearch(target_map []int, value int) int {
	start_index := 0
	end_index := len(target_map) / 2

	for start_index <= end_index {
		mid := (start_index + end_index)
		if target_map[mid] < value {
			start_index = mid + 1
		} else {
			end_index = mid - 1
		}

	}
	if start_index == len(target_map) || target_map[start_index] != value {
		return -1
	}
	return start_index
}

func main() {
	v := []int{1, 4, 5, 8, 9, 7, 12}
	sort.Ints(v)
	fmt.Println(BinarySearch(sort.IntSlice(v), 7))
	fmt.Println(v[BinarySearch(sort.IntSlice(v), 7)])

}
