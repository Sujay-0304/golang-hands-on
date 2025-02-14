package main

import (
	"fmt"
	"strings"
)

func main() {
	// v := 5
	// fmt.Println(v)

	// read()
	// splits()
	// fmt.Println(iterate(1, 10))
	array_operations()
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
		fmt.Println(string(count[i]))
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