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
	slice_operations()  
	// map_operations() 
	// switch_case()
	// i,j := climbing_stairs()
	// fmt.Print(i , j) 
	// types()  
	
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

	if temp := -10; temp >= 0{
		fmt.Println("the value of temp is positive")
	} else {
		fmt.Println("the value of temp is not positive")
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

	// i := 0
	// for i < sec {
	// 	fmt.Print(i, " ")
	// 	i++
	// }

	count := "strings"
	// for i := 0; i < len(count); i++ {
	// 	fmt.Println(string(count[i])) //if we dont convert it to string it will print ascii value
	// }

	for ind,val :=  range count {
		fmt.Printf(" %d %c\n", ind, val) 
		// fmt.Println(string(val)) 
	}

	return 0
}

func slice_operations(){
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[2:3] 
	// fmt.Println(slice) 
	// second_arr := []int{8,9,0}
	third_arr := make([]int, 5) // 3  cap 5  1 1 1  (3) 1 1 1 1 (8)

	third_arr = append(third_arr,1,2,3,4,5,4) 
	third_arr[2] = 7
	fmt.Println(third_arr) 
	fmt.Println(len(third_arr))
	fmt.Println(cap(third_arr))  

	// fmt.Println(arr)
	// fmt.Println(arr[0])
	// fmt.Println(arr[1:3])
	// fmt.Println(arr[:3])
	// fmt.Println(arr[2:])
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))
	// val := append(arr,second_arr...) //spread operator 
	// fmt.Println(val)
	// slice2 := make([]int, 5) 
	// slice2 = append(slice2, 9,8) 
	// fmt.Println(slice2) 
	// slice2[1] = 7 
	// fmt.Println(slice2) 

}

func climbing_stairs() (int, int) {
	val := 5
	arr := make([]int, val+1)
	arr[val-1]  = 1;
	for i:=val-2; i>=0; i--{
		arr[i] = arr[i+1] + arr[i+2] 
	}
	// fmt.Println(arr)
	return arr[0] + arr[1], 0 
}

func types(){ 
	type Age int 
	var myage Age  = 9
	fmt.Print(myage) 

}

func map_operations(){
	m := make(map[string]int) // empty map 
	fmt.Println(m) 
	m["a"] = 1 
	m["b"] = 2 
	m["a"] = 3
	fmt.Println(m) 

	map2 := map[string]int{"a":1} 
	fmt.Println(map2) 
	
	val,exist := m["d"]
	val2,exist2 := m["a"]

	fmt.Println(val,exist) 
	fmt.Println(val2,exist2) 

	for key, value := range m {
		fmt.Println(key, value)
	}

}