package main

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func getTextRaw(node ast.Node) []string {
	switch node := node.(type) {
	// TODO: Add case for *ast.Table
	case *ast.TableHeader:
		raw := node.GetChildren()[0]
		return getTextRaw(raw)
	case *ast.TableBody:
		raw := node.GetChildren()[0]
		return getTextRaw(raw)
	case *ast.TableRow:
		var ss []string
		for _, c := range node.GetChildren() {
			leaf := c.GetChildren()[0].AsLeaf()
			ss = append(ss, string(leaf.Literal))
		}
		return ss
	default:
		return []string{}
	}
}

func getAllTableCell(node ast.Node) [][]string {
	var data [][]string
	for _, child := range node.GetChildren() {
		for _, c := range child.GetChildren() {
			data = append(data, getTextRaw(c))
		}
	}
	return data
}

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
	getAllTableCell(output)
}
