package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func main() {

	os.Clearenv()
	input, err := os.Open(".config")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()


	parser := syntax.NewParser()
	file, err := parser.Parse(input, "")
	if err != nil {
		log.Fatal(err)
	}
	printer := syntax.NewPrinter()

	syntax.Walk(file, func(node syntax.Node) bool {

		switch x := node.(type) {
		case *syntax.DeclClause:

			if x.Variant.Value == "export" {

				var b bytes.Buffer

				name := x.Args[0].Name.Value
				
				printer.Print(&b, x.Args[0].Value)

				if b.Len() > 0 {

					trimmed := strings.Trim(b.String(), "\"")

					fmt.Println(x.Args[0].Name.Value, " =", trimmed)

					v := os.ExpandEnv(trimmed)

					err := os.Setenv(name, v)
					if err != nil {
						log.Fatal(err)
					}
				}

			}

		}

		return true
	})

	fmt.Println("API_MYSQL : ", os.Getenv("API_MYSQL"))
}
