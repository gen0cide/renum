package renum

import (
	"encoding"
	"encoding/json"
	"fmt"
)

// Enum forms the basis for a strongly typed enum class that allows for
// good cross-package interoperability. This creates enums that play nice
// with things like loggers and metrics emitters.
type Enum interface {
	// Coder requires enums to be able to describe their underlying integer representation.
	Coder

	// Namespacer requires enums to be uniquely identifiable with namespace and path values.
	Namespacer

	// Typer requires enums to describe their type semantics relating to source code.
	Typer

	// Descriptioner requires enums to describe themselves in detail, upon request.
	Descriptioner

	// Caser requires enums to describe their names in multiple
	// string case semantics.
	Caser

	// Stringer implements fmt.Print handling
	fmt.Stringer

	// Marshaler implements JSON typing
	json.Marshaler

	// TextMarshaler implements text marshaling
	encoding.TextMarshaler
}

var undefinedValue = `undefined_enum_value`

// IsUndefined is used to check if an enum value is undefined.
func IsUndefined(e Enum) bool {
	if e.String() == undefinedValue {
		return true
	}

	if e.Code() == 0 {
		return true
	}

	return false
}
