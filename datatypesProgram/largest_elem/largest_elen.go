package main

import (
	"fmt"
)

func main() {
	slice := []int{54,34,23,89,67,4,76}
	largest := slice[0] 
	for i,val := range slice {
		fmt.Println(i,val,largest) 
		if val > largest {
			largest = val
		}
	}
	fmt.Println("Largest element is: ",largest)
}