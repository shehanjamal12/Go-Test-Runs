package main

import (
	"fmt"
	"sync"
)
	var wg = sync.WaitGroup{}
func main() {
	var msg string ="hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	
	wg.Add(1)
	i:=1
	j:=2
	go func(i int, j int){
		k:=i+j	
		fmt.Println(k)
		wg.Done()
	}(i,j)
	wg.Wait()
}