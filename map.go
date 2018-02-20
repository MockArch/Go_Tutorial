package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["m"] = 17
	m["o"] = 13
	m["c"] = 12
	m["k"] = 19

	fmt.Println("map:", m)

	fmt.Println("map_len:", len(m))

	delete(m, "m")
	fmt.Println("map:", m)
}