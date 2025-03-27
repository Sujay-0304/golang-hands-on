package main

import (
	"fmt"
	"sync"
)

func countWithoutGoroutine() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func countWithRoutines(wg *sync.WaitGroup, num int) {
	// defer wg.Done()
	defer wg.Done()
	fmt.Println(num)
}

func main() {
	// countWithoutGoroutine()

	// var wg sync.WaitGroup
	// for i := 1; i <= 100; i++ {
	// 	wg.Add(2)
	// 	go countWithRoutines(&wg, i)
	// }
	// wg.Wait() // Wait for all goroutines to finish

	// ch := make(chan int,2) 

	// ch <- 1
	// ch <- 2

	// fmt.Println(<-ch) 
	// fmt.Println(<-ch) 
	// // fmt.Println(<-ch) 
	// fmt.Println("this is the main")
	closing()
	
}
func closing(){
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	// close(ch) 
	ch <- 30
	for val := range ch { 
		fmt.Println(val)
	}
}