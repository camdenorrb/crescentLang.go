package main

import (
	"fmt"
	"sync"
	"unicode"
)

func main() {

	fmt.Println(unicode.IsNumber('1'))

	var wg sync.WaitGroup
	wg.Add(1)
	//fmt.Println(common.Visibility("Meow"))
	c := make(chan string)

	go func() {
		thing(c)
		wg.Done()
	}()
	c <- "Meow"
	c <- "Meow"
	close(c)

	wg.Wait()
	//

	/*
		thing := token.Operator("")

		if _, ok := thing.(token.Operator); ok {

		}*/
}

func thing(values chan string) {
	for value := range values {
		fmt.Println(value)
	}
}
