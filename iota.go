package main

import (
	"fmt"
)

//iota
const (
	a = iota
	b
	c
)
const (
	//iota is specified to each block
	d = iota
	e
	f
)

func main() {

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

}
