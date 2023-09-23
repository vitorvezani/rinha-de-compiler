package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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

	file, err := os.Create("/app/index.js")
	if err != nil {
		log.Fatalf("error creating output file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(code)
	if err != nil {
		log.Fatalf("error writing JS code to file: %v", err)
	}

	nodeCmd := exec.Command("node", "/app/index.js")

	output, err := nodeCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error running node file: %v", err)
	}
	fmt.Print(string(output))
}
