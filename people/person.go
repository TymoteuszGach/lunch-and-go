package people

import "fmt"

type Person struct{
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hello I'm %s", p.name)
}