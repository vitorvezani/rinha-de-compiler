package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var functionParamsCount = map[string]int{}

type AST struct {
	Name       string
	Expression Term
	Location   Location
}

func (p *AST) Visit() (string, error) {
	return p.Expression.Visit()
}

type Location struct {
	Start    int
	End      int
	Filename string
}

type Parameter struct {
	Text     string
	Location Location
}

type Var struct {
	Text     string
	Location Location
}

func (i *Var) Visit() (string, error) {
	return i.Text, nil
}

type Function struct {
	Parameters []Parameter
	Value      Term
	Location   Location
}

func (i *Function) Visit() (string, error) {
	var params []string
	for _, p := range i.Parameters {
		params = append(params, p.Text)
	}
	v, err := i.Value.Visit()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("function(%s) {%s};", strings.Join(params, ", "), v), nil
}

type Visitable interface {
	Visit() (string, error)
}

type Term struct {
	Kind string
	v    Visitable
}

func (t *Term) Visit() (string, error) {
	if t.v != nil {
		return t.v.Visit()
	}
	return "", nil
}

func (t *Term) UnmarshalJSON(b []byte) error {
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	v, ok := m["kind"]
	if !ok {
		return errors.New("missing kind")
	}

	t.Kind = v.(string)

	switch v.(string) {
	case "Int":
		var i Int
		if err := json.Unmarshal(b, &i); err != nil {
			return err
		}
		t.v = &i
	case "Str":
		var s Str
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		t.v = &s
	case "Call":
		var c Call
		if err := json.Unmarshal(b, &c); err != nil {
			return err
		}
		t.v = &c
	case "Binary":
		var bin Binary
		if err := json.Unmarshal(b, &bin); err != nil {
			return err
		}
		t.v = &bin
	case "Function":
		var f Function
		if err := json.Unmarshal(b, &f); err != nil {
			return err
		}
		t.v = &f
	case "Let":
		var l Let
		if err := json.Unmarshal(b, &l); err != nil {
			return err
		}
		t.v = &l
		trackFuntionParamsCount(l)
	case "If":
		var i If
		if err := json.Unmarshal(b, &i); err != nil {
			return err
		}
		t.v = &i
	case "Print":
		var p Print
		if err := json.Unmarshal(b, &p); err != nil {
			return err
		}
		t.v = &p
	case "First":
		var f First
		if err := json.Unmarshal(b, &f); err != nil {
			return err
		}
		t.v = &f
	case "Second":
		var s Second
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		t.v = &s
	case "Bool":
		var boolean Bool
		if err := json.Unmarshal(b, &boolean); err != nil {
			return err
		}
		t.v = &boolean
	case "Tuple":
		var tuple Tuple
		if err := json.Unmarshal(b, &tuple); err != nil {
			return err
		}
		t.v = &tuple
	case "Var":
		var v Var
		if err := json.Unmarshal(b, &v); err != nil {
			return err
		}
		t.v = &v
	default:
		return fmt.Errorf("unknown kind: %v", v)
	}

	return nil
}

type Call struct {
	Callee    Term
	Arguments []Term
	Location  Location
}

func (i *Call) Visit() (string, error) {
	c, err := i.Callee.Visit()
	if err != nil {
		return "", err
	}

	var args []string
	for _, a := range i.Arguments {
		arg, err := a.Visit()
		if err != nil {
			return "", err
		}
		args = append(args, arg)
	}

	if expected, found := functionParamsCount[c]; found {
		if expected != len(args) {
			return "", fmt.Errorf("function %s expects %d arguments, got %d", c, expected, len(args))
		}
	}

	return fmt.Sprintf("%s(%s)", c, strings.Join(args, ", ")), nil
}

type Let struct {
	Name     Parameter
	Value    Term
	Next     Term
	Location Location
}

func trackFuntionParamsCount(l Let) {
	if fn, ok := l.Value.v.(*Function); ok {
		functionParamsCount[l.Name.Text] = len(fn.Parameters)
	}
}

func (i *Let) Visit() (string, error) {
	i.Name.Text = fmt.Sprintf("var %s", i.Name.Text)
	v, err := i.Value.Visit()
	if err != nil {
		return "", err
	}
	n, err := i.Next.Visit()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s = %s; %s", i.Name.Text, v, n), nil
}

type Str struct {
	Value    string
	Location Location
}

func (i *Str) Visit() (string, error) {
	return fmt.Sprintf("'%s'", i.Value), nil
}

type Int struct {
	Value    int
	Location Location
}

func (i *Int) Visit() (string, error) {
	return fmt.Sprintf("%d", i.Value), nil
}

type Bool struct {
	Value    bool
	Location Location
}

func (i *Bool) Visit() (string, error) {
	return fmt.Sprintf("%v", i.Value), nil
}

type If struct {
	Condition Term
	Then      Term
	Otherwise Term
	Location  Location
}

func (i *If) Visit() (string, error) {
	c, err := i.Condition.Visit()
	if err != nil {
		return "", err
	}
	t, err := i.Then.Visit()
	if err != nil {
		return "", err
	}

	o, err := i.Otherwise.Visit()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("if (%s) { return %s; } else { return %s; }", c, t, o), nil
}

type Binary struct {
	LHS      Term
	OP       BinaryOP
	RHS      Term
	Location Location
}

func (i *Binary) Visit() (string, error) {
	lhs, err := i.LHS.Visit()
	if err != nil {
		return "", err
	}

	rhs, err := i.RHS.Visit()
	if err != nil {
		return "", err
	}

	op, err := i.OP.apply()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s %s", lhs, op, rhs), nil
}

type Tuple struct {
	First    Term
	Second   Term
	Location Location
}

func (i *Tuple) Visit() (string, error) {
	f, err := i.First.Visit()
	if err != nil {
		return "", err
	}

	s, err := i.Second.Visit()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("'(%s, %s)'", f, s), nil
}

type First struct {
	Value    Term
	Location Location
}

func (i *First) Visit() (string, error) {
	t, ok := i.Value.v.(*Tuple)
	if !ok {
		return "", errors.New("first expects a tuple")
	}

	f, err := t.First.Visit()
	if err != nil {
		return "", err
	}

	return f, nil
}

type Second struct {
	Value    Term
	Location Location
}

func (i *Second) Visit() (string, error) {
	t, ok := i.Value.v.(*Tuple)
	if !ok {
		return "", errors.New("second expects a tuple")
	}

	f, err := t.Second.Visit()
	if err != nil {
		return "", err
	}

	return f, nil
}

type Print struct {
	Value    Term
	Location Location
}

func (i *Print) Visit() (string, error) {
	v, err := i.Value.Visit()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("console.log(%s);", v), nil
}

type BinaryOP string

const (
	Add BinaryOP = "Add" // Soma	3 + 5 = 8, "a" + 2 = "a2", 2 + "a" = "2a", "a" + "b" = "ab"
	Sub BinaryOP = "Sub" // Subtração	0 - 1 = -1
	Mul BinaryOP = "Mul" // Multiplicação	2 * 2 = 4
	Div BinaryOP = "Div" // Divisão	3 / 2 = 1
	Rem BinaryOP = "Rem" // Resto da divisão	4 % 2 = 0
	Eq  BinaryOP = "Eq"  // Igualdade	"a" == "a", 2 == 1 + 1, true == true
	Neq BinaryOP = "Neq" // Diferente	"a" != "b", 3 != 1 + 1, true != false
	Lt  BinaryOP = "Lt"  // Menor	1 < 2
	Gt  BinaryOP = "Gt"  // Maior	2 > 3
	Lte BinaryOP = "Lte" // Menor ou igual	1 <= 2
	Gte BinaryOP = "Gte" // Maior ou igual	1 >= 2
	And BinaryOP = "And" // Conjunção	true && false
	Or  BinaryOP = "Or"  // Disjunção	false || true
)

func (op BinaryOP) apply() (string, error) {
	switch op {
	case Add:
		return "+", nil
	case Sub:
		return "-", nil
	case Mul:
		return "*", nil
	case Div:
		return "/", nil
	case Rem:
		return "%", nil
	case Eq:
		return "==", nil
	case Neq:
		return "!=", nil
	case Lt:
		return "<", nil
	case Gt:
		return ">", nil
	case Lte:
		return "<=", nil
	case Gte:
		return ">=", nil
	case And:
		return "&&", nil
	case Or:
		return "||", nil
	default:
		return "", fmt.Errorf("unknown binary operator: %v", op)
	}
}
