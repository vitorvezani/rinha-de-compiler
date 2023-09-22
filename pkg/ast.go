package pkg

type File struct {
	name       string
	expression Term
	location   Location
}

type Location struct {
	start    int
	end      int
	filename string
}

type Parameter struct {
	text     string
	Location Location
}

type Var struct {
	kind     string
	text     string
	location Location
}

type Function struct {
	kind       string
	parameters []Parameter
	value      Term
	location   Location
}

type Term interface{ marker() }

func (Int) marker()
func (Str) marker()
func (Call) marker()
func (Binary) marker()
func (Function) marker()
func (Let) marker()
func (If) marker()
func (Print) marker()
func (First) marker()
func (Second) marker()
func (Bool) marker()
func (Tuple) marker()
func (Var) marker()

type Call struct {
	kind      string
	callee    Term
	arguments []Term
	location  Location
}

type Let struct {
	kind     string
	name     Parameter
	value    Term
	next     Term
	location Location
}

type Str struct {
	kind     string
	value    string
	location Location
}

type Int struct {
	kind     string
	value    int
	location Location
}

type Bool struct {
	kind     string
	value    bool
	location Location
}

type If struct {
	kind      string
	condition Term
	then      Term
	otherwise Term
	location  Location
}

type Binary struct {
	kind     string
	lhs      Term
	op       BinaryOP
	rhs      Term
	location Location
}

type Tuple struct {
	kind     string
	first    Term
	second   Term
	location Location
}

type First struct {
	kind     string
	value    Term
	location Location
}

type Second struct {
	kind     string
	value    Term
	location Location
}

type Print struct {
	kind     string
	value    Term
	location Location
}

type BinaryOP int

const (
	Add BinaryOP = iota // Soma	3 + 5 = 8, "a" + 2 = "a2", 2 + "a" = "2a", "a" + "b" = "ab"
	Sub                 // Subtração	0 - 1 = -1
	Mul                 // Multiplicação	2 * 2 = 4
	Div                 // Divisão	3 / 2 = 1
	Rem                 // Resto da divisão	4 % 2 = 0
	Eq                  // Igualdade	"a" == "a", 2 == 1 + 1, true == true
	Neq                 // Diferente	"a" != "b", 3 != 1 + 1, true != false
	Lt                  // Menor	1 < 2
	Gt                  // Maior	2 > 3
	Lte                 // Menor ou igual	1 <= 2
	Gte                 // Maior ou igual	1 >= 2
	And                 // Conjunção	true && false
	Or                  // Disjunção	false || true
)
