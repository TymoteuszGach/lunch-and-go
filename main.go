package main

import (
	"fmt"
	"time"
)

func main() {
	result1 := sum(3, 7)
	a, b := getValues()
	c, err := tryCalculate()

	fmt.Printf("sum: %d\n", result1)
	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Printf("c: %d, err: %v\n", c, err)
}

func sum(a, b int) int {
	return a + b
}

func getValues() (int, int) {
	return 5, 8
}

func tryCalculate() (int, error) {
	if time.Now().Weekday() == time.Friday {
		return 0, fmt.Errorf("chill out")
	}
	return 10, nil
}
