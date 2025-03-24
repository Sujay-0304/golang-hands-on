package main

import (
	"encoding/json"
	"fmt"
	_ "io"
	"os"
)


func main(){
	bytes,err := os.ReadFile("encoder.json") 
	if err != nil {
		fmt.Println("error opening the file ")
		return 
	}
	fmt.Println(bytes) 
	keyword := "BookName" 

	search(bytes, keyword) 


}
func search(data []byte, keyword string){ 
	var rawData map[string]json.RawMessage 
	// var results []string
	fmt.Println(rawData)

	err := json.Unmarshal(data, &rawData)
	if err != nil {
		fmt.Println("cant unmarshal") 
	} 

	fmt.Println(rawData)

}