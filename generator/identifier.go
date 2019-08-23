package generator

import "github.com/gobuffalo/flect"

// Identifier is used to create various cases and pluralities of the enums being generated.
type Identifier flect.Ident

// NewIdentifier creates a new identifier from a provided string.
func NewIdentifier(s string) Identifier {
	return Identifier(flect.New(s))
}

// Snake returns the identifier's snake_case representation.
func (i Identifier) Snake() string {
	return i.Ident().Underscore().String()
}

// Pascal returns the identifiers PascalCase representation.
func (i Identifier) Pascal() string {
	return i.Ident().Pascalize().String()
}

// Screaming returns the identifiers SCREAMING_CASE representation.
func (i Identifier) Screaming() string {
	return i.Ident().Underscore().ToUpper().String()
}

// Ident returns the identifiers underlying flect structure.
func (i Identifier) Ident() flect.Ident {
	return flect.Ident(i)
}
