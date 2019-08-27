package renum

import (
	"database/sql"
	"database/sql/driver"

	"github.com/go-yaml/yaml"
)

// YAMLSerializer is an extension to the Enum type that allows for compound interface types
// where an Enum requires YAML interoperability.
type YAMLSerializer interface {
	yaml.Marshaler
	yaml.Unmarshaler
}

// DBSerializer is an extension to the Enum type that allows for compound interface types
// where an Enum requires database/sql interoperability.
type DBSerializer interface {
	sql.Scanner
	driver.Value
}
