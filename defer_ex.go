package main

import (
	"fmt"
)

func main() {
	function1()
}

func function1() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Printf("this is the fuction1 top")
	defer function2()
	fmt.Printf("/n This is fuction1 down")
}

func function2() {
	fmt.Printf("/n Function2: Deferred until the end of the calling function!")
}
