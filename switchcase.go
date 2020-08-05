package main

import (
	"fmt"
)

func main() {

	switch 2 { //after swirch the case numer must be entered
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("default")
	}
	switch 1 { //after swirch the case numer must be entered
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("default")
	}
	switch 8 { //after swirch the case numer must be entered
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("default")
	}
	var i int = 2
	switch i { //after swirch the case numer must be entered
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("default")
	}

}
