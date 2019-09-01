package renum

// Coder allows an enum to retrieve it's builtin underlying numeric value.
type Coder interface {
	Code() int
}
