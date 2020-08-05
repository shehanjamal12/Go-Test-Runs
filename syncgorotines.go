package main

import (
	"fmt"
	"sync"
)
	var m =sync.RWMutex{}
	var k = 5
	var kg = sync.WaitGroup{}

func main() {
	for i:=0;i<10;i++{
		kg.Add(2)
		m.RLock()
		go increment()
		m.Lock()
		go adding()
	}
	kg.Done()
}
func adding(){
	var j int=k+k
	fmt.Println(j)
	m.Unlock()
} 
func increment(){
	k++
	m.RUnlock()
}