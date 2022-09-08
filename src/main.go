package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {

	/*
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
	*/

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Error("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	) //

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
