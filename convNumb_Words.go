package main 


import (
	"fmt"
	//"log"
)



func main() {
	var response int
	fmt.Println("Enter Number : ")
	_, err := fmt.Scanf("%d", &response)
	if err != nil{
		fmt.Println("err")
}
	if len(response) >4 && len(response) <= 0{
		fmt.Println("Opps!! Try only 4 or less digit ")
	}
	convert_to_words(&response)
}


func convert_to_words (num *string) {
	s := len(*num)
		fmt.Println(s)
	}