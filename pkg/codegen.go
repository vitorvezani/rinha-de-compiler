package pkg

func Codegen(program *AST) (string, error) {
	return program.Visit()
}
