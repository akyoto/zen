package spec

// Operators defines the operators used in the language.
// The number corresponds to the operator priority and can not be zero.
var Operators = map[string]int{
	",": 1,

	"=":   2,
	"+=":  2,
	"-=":  2,
	"*=":  2,
	"/=":  2,
	">>=": 2,
	"<<=": 2,

	"||": 3,

	"&&": 4,

	"==": 5,
	"!=": 5,
	"<=": 5,
	">=": 5,

	"<": 6,
	">": 6,

	"+": 7,
	"-": 7,

	"*": 8,
	"/": 8,
	"%": 8,
}
