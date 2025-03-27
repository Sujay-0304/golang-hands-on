
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("fullscan.json")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}

	keyword := "1801"
	results := search(bytes, keyword)

	if len(results) != 0 {
		fmt.Println("_____")
		fmt.Println(results)
	} else {
		fmt.Println("_____")
		fmt.Println("No result")
	}
}

func search(data []byte, keyword string) []string {
	var results []string
	// fmt.Println(data) 

	var rawData map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err == nil {
		for k, v := range rawData {
			if k == keyword {
				var val string
				json.Unmarshal(v, &val)
				results = append(results, val)
			} else {
		
				results = append(results, search(v, keyword)...)
			}
		}
		return results
	}

	var rawArray []json.RawMessage
	if err := json.Unmarshal(data, &rawArray); err == nil {
		// fmt.Println(rawArray) 
		for _, v := range rawArray {
			results = append(results, search(v, keyword)...)
		}
		return results
	}
	return results
}


// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	_ "io"
// 	"os"
// )


// func main(){
// 	bytes,err := os.ReadFile("encoder.json") 
// 	if err != nil {
// 		fmt.Println("error opening the file ")
// 		return 
// 	}
// 	// fmt.Println(bytes) 
// 	keyword := "BookName" 

// 	search(bytes, keyword) 


// }
// func search(data []byte, keyword string) [] string { 
// 	var rawData map[string]json.RawMessage 
// 	var results []string
// 	// fmt.Println(rawData)

// 	err := json.Unmarshal(data, &rawData)
// 	if err != nil {
// 		fmt.Println("cant unmarshal") 
// 		fmt.Println(err) 
// 	} 
// 	// fmt.Println(rawData)

// 	for k,v := range rawData {
// 		if k == keyword {
// 			var val string
// 			json.Unmarshal(v, &val) 
// 			results = append(results,val) 
// 		} else {
// 			var nestedData interface{} 
// 			json.Unmarshal(v,&nestedData) 
// 			fmt.Println(nestedData)
// 			nestedBytes,_ := json.Marshal(nestedData)
// 			fmt.Println(nestedBytes) 
// 			fmt.Println("**********")
// 			results = append(results, search(nestedBytes,keyword)...) 
// 		}
// 	}
// 	return results 

// }
