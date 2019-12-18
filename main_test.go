package main

import (
	"testing"

	"github.com/gomarkdown/markdown/parser"
)

func TestSampleTable(t *testing.T) {
	input := `
| foo | bar | boo |
|-----|-----|-----|
|  1  |  3  |  5  |
|  2  |  4  |  6  |`
	inputByte := []byte(input)
	parser := parser.New()
	output := parser.Parse(inputByte)
	results := extractTextFromTableDocument(output)
	expected := [][]string{
		{"foo", "bar", "boo"},
		{"1", "3", "5"},
		{"2", "4", "6"}}
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if results[i][j] != expected[i][j] {
				t.Errorf("results[%d][%d] is expected %s, but got=%s", i, j, results[i][j], expected[i][j])
			}
		}
	}
}

func TessContainText(t *testing.T) {
	input := `
This node is unrelated.

| foo | bar | boo |
|-----|-----|-----|
|  1  |  3  |  5  |
|  2  |  4  |  6  |`
	inputByte := []byte(input)
	parser := parser.New()
	output := parser.Parse(inputByte)
	results := extractTextFromTableDocument(output)
	expected := [][]string{
		{"foo", "bar", "boo"},
		{"1", "3", "5"},
		{"2", "4", "6"}}
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if results[i][j] != expected[i][j] {
				t.Errorf("results[%d][%d] is expected %s, but got=%s", i, j, results[i][j], expected[i][j])
			}
		}
	}
}
