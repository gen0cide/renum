package generator

import (
	"strings"

	"github.com/go-openapi/swag"

	"github.com/gen0cide/flect"
)

var acronyms = []string{
	"AMI",
}

func init() {
	swag.AddInitialisms(acronyms...)
}

// AddIdentifierInitialism is used to add custom initializations into the library.
func AddIdentifierInitialism(s ...string) {
	swag.AddInitialisms(s...)
}

// Identifier is used to create various cases and pluralities of the enums being generated.
type Identifier struct {
	original string
}

// NewIdentifier creates a new identifier from a provided string.
func NewIdentifier(s string) Identifier {
	return Identifier{
		original: s,
	}
}

// Snake returns the identifier's snake_case representation.
func (i Identifier) Snake() string {
	return swag.ToFileName(i.original)
}

// Pascal returns the identifiers PascalCase representation.
func (i Identifier) Pascal() string {
	return swag.ToGoName(i.original)
}

// Screaming returns the identifiers SCREAMING_CASE representation.
func (i Identifier) Screaming() string {
	return strings.ToUpper(i.Snake())
}

// Ident returns the identifiers underlying flect structure.
func (i Identifier) Ident() flect.Ident {
	return flect.New(i.original)
}
