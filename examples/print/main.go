package main

import (
	"fmt"
	"log"

	conllu "github.com/nuvi/go-conllu"
)

func main() {
	sentences, err := conllu.ParseFile("../../test_data/en_ewt-ud-train.small.conllu")
	if err != nil {
		log.Fatal(err)
	}

	for _, sentence := range sentences {
		for _, token := range sentence.Tokens {
			fmt.Println(token)
		}
		fmt.Println()
	}
}
