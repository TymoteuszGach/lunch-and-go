package main

import "fmt"

func main() {
	i := 1
	var j *int
	j = &i
	fmt.Printf("i: %d\n", i)
	fmt.Printf("j: %d\n", *j)

	zeroval(i)
	fmt.Printf("zeroval: %d\n", i)

	zeroptr(&i)
	fmt.Printf("zeroptr: %d\n", i)
}

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}