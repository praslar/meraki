package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var a []string

func main() {
	fmt.Println("hello")
	a = []string{"a", "b", "c", "d"}
	for _, i := range a {
		wg.Add(1)
		go printSomething(i)
	}
	wg.Wait()
}

func printSomething(i string) {
	fmt.Println(i)
	wg.Done()
}
