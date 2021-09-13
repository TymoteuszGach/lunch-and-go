package main

import (
	"fmt"
	"github.com/tymoteuszgach/lunch-and-go/animals"
	"github.com/tymoteuszgach/lunch-and-go/people"
)

type greeter interface{
	Greet() string
}

func main() {
	greeters := getGreeters()

	for _, greeter := range greeters {
		greeting := greeter.Greet()
		fmt.Printf("%T greets: %s\n", greeter, greeting)
	}
}

func getGreeters() []greeter {
	return []greeter{
		animals.NewDog(),
		people.NewPerson("Julian Tuwim"),
	}
}
