package main


import "fmt"


func main() {
	message := make(chan string, 2)

	message <- "mock"
	message <- "Arch"
	// message <- "Hell"

	fmt.Println(<-message)
	fmt.Println(<-message)
	// fmt.Println(<-message)

}