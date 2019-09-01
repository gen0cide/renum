package renum

import (
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

	// Sourcer requires enums to be able to self describe aspects of the Go source and package
	// which they're located. This makes Enum's great for tracing and error handling.
	Sourcer

	// Typer requires enums to describe their type.
	Typer

	// Descriptioner requires enums to describe themselves in detail, upon request.
	Descriptioner

	// Caser requires enums to describe their names in multiple
	// string case semantics.
	Caser

	// Stringer implements fmt.Print handling
	fmt.Stringer

	// Marshaler requires that enums be able to support encoding/decoding for
	// a variety of common formats. The expectation is that if your Enum implements
	// Marshaler, that it will also implement the pointer recievers for Unmarshaler.
	// If you're generating your Enum with the Renum CLI, this will happen
	// automatically for you.
	Marshaler
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
