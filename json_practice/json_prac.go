package  main

import (
	// "fmt"
	// "encoding/json"
)

type Person struct {
	NaMe string `json:name`
	Age int `json:"age"` 
	Gender string `json:"gender"` 
}

// func main(){
// 	p := Person {
// 		NaMe: "John",
// 		Age: 25,
// 		Gender: "M", 
// 	}
// 	fmt.Println(p)
// 	data,err := json.MarshalIndent(p,"","   ")	
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(data))
// }