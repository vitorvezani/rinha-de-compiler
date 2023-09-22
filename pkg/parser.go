package pkg

import "encoding/json"

func Parse(body []byte) (*AST, error) {
	var program AST
	err := json.Unmarshal(body, &program)
	if err != nil {
		return nil, err
	}
	return &program, nil
}
