package main

import (
	"os"

	"github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/example002.bb")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Print()
	}
}
