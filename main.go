package main

import (
	"fmt"
	"time"
)

func sum(n int, result chan<- int){
	sum := 0
	for i:= 1; i <= n; i++{
		sum += i
	}

	time.Sleep(1*time.Second)

	result <- sum
}

func main() {
	resultChan := make(chan int)
	go sum(10, resultChan)

	doSomethingElse()

	select{
	case result := <-resultChan:
		fmt.Printf("result is: %d\n", result)
	case <-time.After(2*time.Second):
		fmt.Println("timeout!!11")
	}
}

func doSomethingElse() {
	fmt.Println("doing something else...")
}
