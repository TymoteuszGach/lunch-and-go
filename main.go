package main

import (
	"fmt"
	"time"
)

func main() {
	run()
}

func run() {
	fmt.Println("Starting task...")
	defer clean()

	time.Sleep(time.Second)
	fmt.Println("Task finished...")
}

func clean() {
	fmt.Println("Cleaning...")
}
