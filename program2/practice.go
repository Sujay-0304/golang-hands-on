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

	// arr := []int{1, 2, 3, 4, 5}
	// array := new ([5]int)
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
