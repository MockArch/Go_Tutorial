package main


import "fmt"


func add1(a *int) int {
	fmt.Printf("&x = ", a)
	*a = *a + 1
	return *a
}


func main() {
	x := 3

	fmt.Printf("x = %d ", x)
	x1 := add1(&x)

	fmt.Printf("x+1 = %d", x1)
	fmt.Printf("x = %d", x)
}