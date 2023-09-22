package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vitorvezani/rinha-de-compiler/pkg"
)

func main() {
	path := os.Args[1:][0]

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	program, err := pkg.Parse(b)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	code, err := pkg.Codegen(program)
	if err != nil {
		log.Fatalf("error in codegen: %v", err)
	}

	fmt.Println(code)

	err = pkg.Eval(code)
	if err != nil {
		log.Fatalf("error evaluating program: %v", err)
	}
}
