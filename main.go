package main

import "fmt"

func main() {
	var name1 string
	name1 = "Andrzej"

	name2 := "Alfred"

	name3 := getName3()

	fmt.Printf("Hello %s!\n", name1)
	fmt.Printf("Hello %s!\n", name2)
	fmt.Printf("Hello %s!\n", name3)
}

func getName3() string {
	return "Antoni"
}