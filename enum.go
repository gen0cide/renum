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

// EnumTypeInfo is a type used to hold all the metadata associated with a given renum.Enum where
// the fields of this structure are associated directly with the return values from the renum.Enum interface.
// This acts as a convenience to help things like structured loggers or HTTP JSON responses to be have
// information extracted into a self contained object.
type EnumTypeInfo struct {
	Name        string `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Code        int    `json:"code,omitempty" mapstructure:"code,omitempty" yaml:"code,omitempty" toml:"code,omitempty"`
	Namespace   string `json:"namespace,omitempty" mapstructure:"namespace,omitempty" yaml:"namespace,omitempty" toml:"namespace,omitempty"`
	Path        string `json:"path,omitempty" mapstructure:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	Kind        string `json:"kind,omitempty" mapstructure:"kind,omitempty" yaml:"kind,omitempty" toml:"kind,omitempty"`
	Source      string `json:"source,omitempty" mapstructure:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty"`
	ImportPath  string `json:"import_path,omitempty" mapstructure:"import_path,omitempty" yaml:"import_path,omitempty" toml:"import_path,omitempty"`
	Description string `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
}

// ExtractEnumTypeInfo is used to take a renum.Enum type and expand it's details into a more annotated
// structure. The primary purpose of this is to act as a helper to loggers who wish to expand interface methods
// of the renum.Enum type into a nested, flat structure.
func ExtractEnumTypeInfo(e Enum) EnumTypeInfo {
	return EnumTypeInfo{
		Name:        e.String(),
		Code:        e.Code(),
		Namespace:   e.Namespace(),
		Path:        e.Path(),
		Kind:        e.Kind(),
		Source:      e.Source(),
		ImportPath:  e.ImportPath(),
		Description: e.Description(),
	}
}
