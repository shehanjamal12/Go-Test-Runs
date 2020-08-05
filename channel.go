package main

import (
	"fmt"
	"sync"
)

	var lg = sync.WaitGroup{}
	
func main(){
	ch := make(chan int,50)

	lg.Add(2)
	go func(ch  <- chan int){
		for i:= range ch{
			fmt.Println(i)
		}
		lg.Done()

	}(ch)
	go func(ch chan <-int){
		ch <-42
		close(ch)
		lg.Done()
	}(ch)

		lg.Wait()

}
