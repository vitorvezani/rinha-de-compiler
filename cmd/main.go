package main

import (
	"log"
	"os"

	"github.com/vitorvezani/rinha-de-compiler/pkg"
)

func main() {
	b, err := os.ReadFile("/var/rinha/source.rinha.json")
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

	err = pkg.Eval(code)
	if err != nil {
		log.Fatalf("error evaluating program on otto JS runtime : %v", err)
	}
}
