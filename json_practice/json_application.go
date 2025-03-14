package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	BookName string `json:"BookName"`
	Title string `json:"Title"` 
	Genre string `json:"Genre"` 
} 

type person struct{
	PersonName string `json:"PersonName"` 
	Gender string `json:"Gender"` 
	Age int `json:"Age"`
}

type AutherBookStore struct {
	Person person `json:"person"`
	Books []Book `json:"Books"`
}
// AutherBookStore -> {
// 	person: {name gender age}, 
// 	books: { 
// 		{bookName,title genre}
// 		{bookName,title genre}
// 		{bookName,title genre}
// 	}
// }
func main(){
	fmt.Println()
	createJson() 
	// loadJSON()

}

func createJson(){
	bookStore := AutherBookStore{
		Person: person{"sujay", "M", 21}, 
		Books: []Book {
			{"1","a","genre1"},
			{"2","b","genre2"},
			{"3","c","genre3"},
			{"4","d","genre4"},
			{"5","e","genre5"},
		}, 
	}
	file,err := json.MarshalIndent(bookStore, "","   ")
	if err != nil {
		fmt.Println("there is error in file")
		return 
	}
	fmt.Println(string(file)) 
	// saveJSON_OS(file)
	saveJSON_ENCODER(bookStore)

	for ind,books := range bookStore.Books{
		fmt.Println(ind + 1) 
		fmt.Println("book genre:", books.Genre) 
	}
}

func saveJSON_OS(data []byte){
	err := os.WriteFile("Data.json",data,0644) 
	if err != nil{
		fmt.Println("there is a problem in creating the file.. in Save JSON")
		return
	}

}
func saveJSON_ENCODER(bookStore AutherBookStore){
	file,err := os.Create("encoder.json")
	defer file.Close() 
	if err != nil{
		return 
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("","  ") 
	encoder.Encode(bookStore) 
	
}

func loadJSON(){ 
	file,err := os.OpenFile("Data.json", os.O_RDONLY , 0644)
	if err != nil {
		fmt.Println("file has error while opening")
	} 
	defer file.Close()

	decoder := json.NewDecoder(file) 
	var bookstore AutherBookStore
	errors := decoder.Decode(&bookstore)
	if errors != nil {
		return 
	}

	fmt.Println("in loadJson")
	// fmt.Println(bookstore) 
	formatedJSON,err := json.MarshalIndent(bookstore,"","    ")
	fmt.Println(string(formatedJSON)) 

}
