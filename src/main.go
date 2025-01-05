package main

import (
	"bufio"
	"crescentLang/common"
	"crescentLang/crescent"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("testdata/hello_world.moon")
	if err != nil {
		panic(err)
	}

	lexer, err := common.NewGenericLexer(crescent.Syntax)
	if err != nil {
		panic(err)
	}

	tokens, err := lexer.Lex(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println(token.Type, token.Value)
	}

	//parsed, err := crescent.Parse(tokens)
	//if err != nil {
	//	panic(err)
	//}

	/*
		for _, node := range parsed {
			fmt.Printf("%+v\n", node)
		}*/

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

	/*
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()
		logger.Error("failed to fetch URL",
			// Structured context as strongly typed Field values.
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)*/ //

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
