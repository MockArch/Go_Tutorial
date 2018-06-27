package main 

import "fmt"

func main() {
	x := Min(1,100, 56, 487, 4238762348, 4874)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{37,326,3523,6587,587,5487,8945,845}
	x = Min(arr...)
	fmt.Printf("The minimum is: %d\n", x)
}


func Min(a...int) (int) {
	if len(a)== 0{
		return 0
	}
	min := a[0]
	for _,v := range a {
		if v < min {
			min = v
		}
	}
	return min
}