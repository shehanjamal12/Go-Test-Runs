package main

import (
	"fmt"
)

func main() {
	var i string
	i = "hey there"
	mess(i)
	i = "yo"
	sum(i, 1, 2, 3, 4, 5, 6, 7, 8)
	var val1 int = 1
	var val2 int = 1
	var j int = sum2(val1, val2)
	fmt.Println(j)

	nmae := name{
		name: "shehan",
	}
	nmae.naming()
}

//function making
func mess(msg string) {
	fmt.Println(msg)
}

func sum(msg string, values ...int) {
	fmt.Println(msg, values)
}

//returning
func sum2(val1 int, val2 int) (val3 int) {
	val3 = val1 + val2
	return
}

//function with struct
type name struct {
	name string
}

func (nmae name) naming() {
	fmt.Println("hey", nmae.name)
}
