package renum

// Caser defines an interface for types (specifically enums) to be able to
// describe their name typically returned by the fmt.Stringer interface
// as various text casings. This is helpful in situations where there are character
// restrictions that are enforced.
type Caser interface {
	// snake_case
	SnakeCase() string

	// PascalCase
	PascalCase() string

	// camelCase
	CamelCase() string

	// SCREAMING_CASE
	ScreamingCase() string

	// command-case
	CommandCase() string
}
