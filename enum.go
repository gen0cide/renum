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
	// Requires enums to be able to describe their underlying integer representation.
	Coder

	// Requires enums to be uniquely identifiable with namespace and path values.
	Namespacer

	// Requires enums to describe their type semantics relating to source code.
	Typer

	// Requires enums to describe themselves in detail, upon request.
	Descriptioner

	// Requires enums to describe their names in multiple
	// string case semantics.
	Caser

	// fmt.Print handling
	fmt.Stringer

	// JSON typing
	json.Marshaler

	// Text
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
