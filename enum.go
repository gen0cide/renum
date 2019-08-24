package renum

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"

	"github.com/go-yaml/yaml"
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
	json.Unmarshaler

	// Text
	encoding.TextMarshaler
	encoding.TextUnmarshaler

	// YAML typing
	yaml.Marshaler
	yaml.Unmarshaler

	// Database I/O
	sql.Scanner
	driver.Value
}

type EnumTypeInfo struct {
	Name      string
	Code      int
	Namespace string
}
