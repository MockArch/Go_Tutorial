package main 

import "fmt"
import "time"


func worker(done chan bool) {
	fmt.Print("\nworking....\n")
	time.Sleep(time.Second)
	fmt.Println("\ndone\n")

	done <- true
}


func main() {
	done := make(chan bool, 1)
	done1 := make(chan bool, 1)

	go worker(done)
	go worker(done1)

	<- done
	<- done1

}