package main 

import "fmt"

func main() {
	var num1 int = 99

	switch num1 {
	case 98, 99:
		fmt.Println("Its equal to 98")

	case 100:
		fmt.Println("its equal to 100")
	default:
		fmt.Println("its not equal to 98 or 100")		
	}
}