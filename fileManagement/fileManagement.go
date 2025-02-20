package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// Read_file()
	// Create_file()
	// Append_file() 
	ReplaceContents()
}

func Read_file(){
	data, err := os.ReadFile("sample.txt"); 
	if err != nil {
		fmt.Println("error reading file")
		return 
	}
	// fmt.Println(data) // data is of type []byte ascii values in list space for 32 'a' for 97 so on.... 
	fmt.Println(string(data))
	// fmt.Println(err)
}

func Create_file(){
	file,err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer file.Close() 
	file.WriteString("Hello World\n")
	file.WriteString("World Hello")
}

func Append_file(){
	file,err :=  os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644) // 0644 is permission
	//  0 -> octal 6 -> read and write permission for owner 4 -> read permission for group and others
	if err != nil {
		fmt.Println("error opening file")
		return
	} 
	defer file.Close()
	bytes, err := file.WriteString("\nyo yo pudichadu nan than\n")
	fmt.Println(err)

	if err != nil {
		fmt.Println("error writing to file")
		return
	}
	fmt.Println(bytes)
}

func ReplaceContents(){
	// text := "hi how are are you doing"
	target := "line"
	replacement := "fine" 
	// newText := strings.ReplaceAll(text, target, replacement) 
	// fmt.Println(newText)
	file,err := os.OpenFile("sample.txt", os.O_WRONLY, 0644) 
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()
	data,err := os.ReadFile("sample.txt") 
	if err != nil {
		fmt.Println("error reading file")
		return 
	}
	newText := strings.ReplaceAll(string(data), target, replacement) 
	file.WriteString(newText) 
}