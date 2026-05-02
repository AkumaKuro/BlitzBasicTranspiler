package main

import (
	"os"

	"github.com/AkumaKuro/BlitzBasicTranspiler/src/lexer"
	"github.com/AkumaKuro/BlitzBasicTranspiler/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/example004.bb")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
