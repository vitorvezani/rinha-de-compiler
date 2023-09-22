package pkg

import "github.com/robertkrimen/otto"

func Eval(code string) error {

	vm := otto.New()
	_, err := vm.Run(code)

	return err
}
