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
	Coder
	Namespacer
	Typer
	Descriptioner

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
