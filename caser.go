package renum

// Caser defines an interface for types (specifically enums) to be able to
// describe their name typically returned by the fmt.Stringer interface
// as various text casings. This is helpful in situations where there are character
// restrictions that are enforced.
type Caser interface {
	// SnakeCase should return enum names formatted as "snake_case" representations
	SnakeCase() string

	// PascalCase should return enum names formatted as "PascalCase" representations
	PascalCase() string

	// CamelCase should return enum names formatted as "camelCase" representations
	CamelCase() string

	// ScreamingCase should return enum names formatted as "SCREAMING_CASE" representations
	ScreamingCase() string

	// CommandCase should return enum names formatted as "command-case" representations
	CommandCase() string

	// TrainCase should return enum names formatted as "TRAIN-CASE" representations
	TrainCase() string

	// DottedCase should renum enum names formatted as "dotted.case" representations
	DottedCase() string
}
