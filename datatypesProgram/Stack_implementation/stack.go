package main

import (
	"fmt"
)

type Stack struct {
	stack []int
}

func (s *Stack) push(val int){
	s.stack  = append(s.stack, val) 
}

func (s *Stack) pop() int {
	if(len(s.stack) == 0){
		return -1 
	}
	val := s.stack[len(s.stack) - 1] 
	s.stack = s.stack[:len(s.stack) - 1]
	return val
}

func (s *Stack) peek() int {
	if(len(s.stack) == 0){
		return -1
	} 
	return s.stack[len(s.stack) - 1]
}

func (s *Stack) displayStack(){
	for i := 0; i < len(s.stack); i++ {
		fmt.Print(s.stack[i], " ")
	}
	fmt.Println()
}


func main(){
	stack := Stack{} 
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.displayStack()
	fmt.Println(stack.pop()) 
	fmt.Println(stack.peek())
	stack.displayStack()
	fmt.Println(stack.pop())
	stack.displayStack() 
}