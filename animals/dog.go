package animals

type Dog struct {}

func NewDog() *Dog {
	return &Dog{}
}

func (d *Dog) Greet() string {
	return "Hau hau"
}