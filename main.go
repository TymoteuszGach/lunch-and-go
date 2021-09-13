package main

import "fmt"

type movie struct {
	title string
	year int
}

func newMovie(title string, year int) *movie {
	return &movie{title: title, year: year}
}

func main() {
	godfather := newMovie("The Godfather", 1972)
	fmt.Println(godfather)

	pulpFiction := movie{title: "Pulp Fiction", year: 1994}
	fmt.Println(pulpFiction)

	fightClub := movie{title: "Fight Club"}
	fmt.Println(fightClub)

	fightClub.year = 1999
	fmt.Println(fightClub)
}
