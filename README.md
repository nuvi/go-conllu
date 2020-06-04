# go-conllu
CoNLL-U parser written in Go. Convert CoNLL-U files to in-memory structs

[![](https://godoc.org/github.com/nuvi/go-conllu?status.svg)](https://godoc.org/github.com/nuvi/go-conllu)

The Computational Natural Language Learning - U format (CoNLL-U) is used by the [Universal Dependencies](https://universaldependencies.org/format.html) project to represent natural language annotations.
`go-conllu` parses CoNNL-U file formats and exposes the data via in-memory Go structs.

## ⚙️ Installation

```bash
go get github.com/nuvi/go-conllu
```

## 🚀 Quick Start

```go
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
```
## Issues

All issues should be submitted via the issues tab on Github. Please provide the code and data used in order for us to reproduce the issue.

## 💬 Contact

Feel free to reach out with questions/comments to maintainers:

[![Twitter Follow](https://img.shields.io/twitter/follow/wagslane.svg?label=Follow%20Wagslane&style=social)](https://twitter.com/intent/follow?screen_name=wagslane)

## Transient Dependencies

None, and we plan to keep it that way.

## 👏 Contributing

We love help! Contribute by forking the repo and opening pull requests. Please ensure that your code passes the existing tests and linting processes, and write new tests to test your changes if applicable.

All pull requests should be submitted to the "master" branch.

```bash
go test
```

```bash
go fmt
```
