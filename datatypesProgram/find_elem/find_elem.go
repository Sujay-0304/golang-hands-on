package main 
import (
	"fmt" 
)
func main(){
	var arr = []int{1,2,3,4,5,6,7,8,9,10} 
	target := 5 
	fmt.Println("The index of the target element is: ")
	fmt.Println(linear(arr, target))
	fmt.Println(binary_search(arr, target))
}

func linear (arr []int, target int ) int {
	ind := -1 
	for i:=0; i<len(arr); i++{
		if arr[i] == target{
			ind = i
		}
	}
	return ind 
}
func binary_search(arr []int, target int) int {
	ind := -1
	low := 0
	high := len(arr) - 1 
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			ind = mid 
			break 
		} else if arr[mid]  < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ind 
}
