package renum

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// PointerUnmarshaler defines how a concrete non-pointer value can conform to the Unmarshaler
// contract by returning a pointer to it's value receiver. This interface has no practical use
// and should not need to be used. It simply exists to allow the Go compiler to ensure compliance
// with the Unmarshaler interface by renum.Enum types.
type PointerUnmarshaler interface {
	PointerUnmarshal() Unmarshaler
}

// Marshaler defines the required methods for a renum.Enum type implementation to properly support
// serialization to various encoding formats (text, json, yaml, sql, flags). This allows for easy
// integration with existing data structures that are commonly serialized into these formats.
type Marshaler interface {
	encoding.TextMarshaler
	json.Marshaler
	yaml.Marshaler
	driver.Valuer
}

// Unmarshaler defines the required methods for a renum.Enum type implementation to properly support
// de-serialization out of various encoding foramts (text, json, yaml, sql, flags). This allows for
// easy integration with existing data structures that are commonly unmarshaled from these formats.
type Unmarshaler interface {
	encoding.TextUnmarshaler
	json.Unmarshaler
	yaml.Unmarshaler
	sql.Scanner
	FlagUnmarshaler
}

// FlagUnmarshaler is used to enforce that enum types can be properly encoded and decoded into
// command line flags without custom implementations. This requires renum.Enums to conform to
// pflag.Value (github.com/spf13/pflag), as well as the standard library flag package.
type FlagUnmarshaler interface {
	String() string
	Set(string) error
	Get() interface{}
	Type() string
}
