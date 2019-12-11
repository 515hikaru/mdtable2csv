package main

import (
	"os"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	input := `
| foo | bar | booo |
|-----|-----|------|
| 1   | 3   | 5    |
| 1   | 9   | 25   |
`
	inputByte := []byte(input)

	parser := parser.New()
	output := parser.Parse(inputByte)
	ast.Print(os.Stdout, output)
}
