package config

import (
	"fmt"
	"path"

	"github.com/pkg/errors"
)

// GoConfig holds the primary configuration information about the Go environment that will be generated into.
type GoConfig struct {
	Type        string     `json:"type,omitempty" mapstructure:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty" default:"int" validate:"required,oneof=int8 int16 int32 int64 int uint uint8 uint16 uint32 uint64"`
	Comment     string     `json:"comment,omitempty" mapstructure:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty" validate:"required"`
	Name        string     `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" validate:"required"`
	Filename    string     `json:"filename,omitempty" mapstructure:"filename,omitempty" yaml:"filename,omitempty" toml:"filename,omitempty"`
	Initialisms []string   `json:"initialisms,omitempty" mapstructure:"initialisms,omitempty" yaml:"initialisms,omitempty" toml:"initialisms,omitempty"`
	PackageName string     `json:"package_name,omitempty" mapstructure:"package_name,omitempty" yaml:"package_name,omitempty" toml:"package_name,omitempty" validate:"required"`
	PackagePath string     `json:"package_path,omitempty" mapstructure:"package_path,omitempty" yaml:"package_path,omitempty" toml:"package_path,omitempty" validate:"required"`
	SkipUnknown bool       `json:"skip_unknown,omitempty" mapstructure:"skip_unknown,omitempty" yaml:"skip_unknown,omitempty" toml:"skip_unknown,omitempty"`
	Prefix      Identifier `json:"-" yaml:"-" toml:"-"`
}

// OutputFilename will either return the user-specified filename, or return a generated filename
// in the form of "generated_%s.go" where %s is the enum type name, lower snake cased, and pluralized.
// An example might be:
//
// 	// Go definition
// 	type ErrCode int
// 	// filename = generated_err_codes.go
//
//
func (g GoConfig) OutputFilename() string {
	if g.Filename != "" {
		return g.Filename
	}
	return fmt.Sprintf("generated_%s.go", g.Prefix.Ident().Underscore().Pluralize().String())
}

// Validate ensures that the package conforms to the stated (package_name is the stated package_path package)
func (g GoConfig) Validate() error {
	if path.Base(g.PackagePath) != g.PackageName {
		return errors.New("specified package_name does not match package_path")
	}

	return nil
}
