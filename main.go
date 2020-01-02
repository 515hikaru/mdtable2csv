package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

const version = "0.1.1"

var (
	toCode      string
	inFile      string
	outFile     string
	showVersion bool
)

func init() {
	flag.StringVar(&toCode, "to-code", "UTF-8", "Use to-encoding for output characters.")
	flag.StringVar(&inFile, "if", "", "Read from file instead of stdin")
	flag.StringVar(&outFile, "of", "", "Write to file instead of stdout")
	flag.BoolVar(&showVersion, "version", false, "Show version")
}

func validateToCode(toCode string) bool {
	codes := []string{"UTF-8", "UTF-8-BOM"}
	for _, code := range codes {
		if strings.ToUpper(toCode) == code {
			return true
		}
	}

	return false
}

func inputFromStdin() string {
	var text string
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		text += scan.Text() + "\n"
	}
	return text
}

func inputFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func extractTextFromChildren(node ast.Node) [][]string {
	var texts [][]string
	for _, raw := range node.GetChildren() {
		rawText := extractTextFromTableDocument(raw)
		for _, text := range rawText {
			texts = append(texts, text)
		}
	}
	return texts
}

func extractTextFromTableDocument(node ast.Node) [][]string {
	switch node := node.(type) {
	case *ast.Document:
		// TODO: Deal with multi tables.
		for _, node := range node.GetChildren() {
			if val := extractTextFromChildren(node); val != nil {
				return val
			}
		}
		return nil
	case *ast.Table:
		texts := extractTextFromChildren(node)
		return texts
	case *ast.TableHeader:
		texts := extractTextFromChildren(node)
		return texts
	case *ast.TableBody:
		texts := extractTextFromChildren(node)
		return texts
	case *ast.TableRow:
		var row [][]string
		var ss []string
		for _, c := range node.GetChildren() {
			if len(c.GetChildren()) == 0 {
				ss = append(ss, "")
				continue
			}
			leaf := c.GetChildren()[0].AsLeaf()
			ss = append(ss, string(leaf.Literal))
		}
		row = append(row, ss)
		return row
	default:
		return nil
	}
}

func dumpCSV(records [][]string, buf *bytes.Buffer) {

	w := csv.NewWriter(buf)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			panic("Write Error")
		}
		w.Flush()
	}
}

func main() {
	flag.Parse()
	if showVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}
	if !validateToCode(toCode) {
		fmt.Printf("%s is an unsupported character code.\n", toCode)
		os.Exit(1)
	}
	var input string
	if inFile != "" {
		var err error
		input, err = inputFromFile(inFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	} else {
		input = inputFromStdin()

	}
	inputByte := []byte(input)
	parser := parser.New()
	output := parser.Parse(inputByte)
	records := extractTextFromTableDocument(output)
	buf := new(bytes.Buffer)
	if strings.ToUpper(toCode) == "UTF-8-BOM" {
		buf.Write([]byte{0xEF, 0xBB, 0xBF})
	}
	dumpCSV(records, buf)
	if outFile == "" {
		fmt.Printf(buf.String())
	} else {
		file, err := os.Create(outFile)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		_, err = file.Write(buf.Bytes())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
