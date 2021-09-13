package main

import (
	"fmt"
	"time"
)

func printNTimes(text string, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("text: %s, i: %d\n", text, i)
	}
}

func main() {
	printNTimes("something", 3)

	go printNTimes("else", 10)

	fmt.Println("some other text")

	time.Sleep(time.Second)
}
