package main

import "fmt"

type movie struct {
	title string
	year int
}

func (m movie) printYear() {
	fmt.Printf("year: %d\n", m.year)
}

func (m *movie) printDescription() {
	fmt.Printf("This is '%s' from %d\n", m.title, m.year)
}

func newMovie(title string, year int) *movie {
	return &movie{title: title, year: year}
}

func main() {
	godfather := newMovie("The Godfather", 1972)

	godfather.printYear()
	godfather.printDescription()

	pulpFiction := movie{title: "Pulp Fiction", year: 1994}

	pulpFiction.printYear()
	pulpFiction.printDescription()

	//var emptyMovie *movie
	//emptyMovie.printDescription()
}
