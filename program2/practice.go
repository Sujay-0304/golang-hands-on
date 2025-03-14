package main

import (
	"fmt"
	"strings"
)

func main() {
	// if_else()
	// read()
	// splits()
	// fmt.Println(iterate(1, 10))
	// array_operations()
	// switch_case()
	// climbing_stairs()
}
func switch_case() {
	val := 10
	switch val {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	default:
		fmt.Println("Value is not 1, 2 or 3")
	}
}

func if_else() {
	val := 10
	if val > 10 {
		fmt.Println("Value is greater than 10")
	} else if val < 10 {
		fmt.Println("Value is less than 10")
	} else {
		fmt.Println("Value is 10")
	}
}

func read() {
	var val int
	fmt.Println("Enter a number: ")
	fmt.Scan(&val)
	fmt.Println(val)

}

func splits() {
	s := "Hello, World!     hos aa     aa a    ksdf    "
	// fmt.Println(s)
	arr := strings.Fields(s)
    // arr = strings.Split(s, " ")
	// arr = strings.Split(string(arr), " ")
	fmt.Println(arr)
	fmt.Println(len(arr))
	// joined := strings.Join(arr, "-")
	// fmt.Println(joined)
}

func iterate(fir int, sec int) int {
	for i := fir; i < sec; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < sec {
		fmt.Print(i, " ")
		i++
	}

	count := "stringa"
	for i := 0; i < len(count); i++ {
		fmt.Println(string(count[i])) //if we dont convert it to string it will print ascii value
	}
	return 0
}

func array_operations(){
	arr := []int{1, 2, 3, 4, 5}
	second_arr := []int{8,9,0}
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1:3])
	fmt.Println(arr[:3])
	fmt.Println(arr[2:])
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	val := append(arr,second_arr...) //spread operator 
	fmt.Println(val)
}

func climbing_stairs() int{
	val := 5
	arr := make([]int, val+1)
	arr[val-1]  = 1;
	for i:=val-2; i>=0; i--{
		arr[i] = arr[i+1] + arr[i+2] 
	}
	fmt.Println(arr)
	return arr[0] + arr[1] 
}