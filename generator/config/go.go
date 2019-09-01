package config

import "github.com/creasty/defaults"

// Go holds the configuration for Go related enum generation settings.
type Go struct {
	Type    *GoType    `json:"type,omitempty" mapstructure:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`
	Build   *GoBuild   `json:"build,omitempty" mapstructure:"build,omitempty" yaml:"build,omitempty" toml:"build,omitempty"`
	Package *GoPackage `json:"package,omitempty" mapstructure:"package,omitempty" yaml:"package,omitempty" toml:"package,omitempty"`
}

// GoType is a sub-configuration for holding information about the enum's Go type alias, and it's declaration.
type GoType struct {
	Numeric     string `json:"numeric,omitempty" mapstructure:"numeric,omitempty" yaml:"numeric,omitempty" toml:"numeric,omitempty" default:"int" validate:"required,oneof=int8 int16 int32 int64 int uint uint8 uint16 uint32 uint64"`
	Name        string `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" validate:"required"`
	Comment     string `json:"comment,omitempty" mapstructure:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty"`
	SkipDeclare bool   `json:"skip_declare,omitempty" mapstructure:"skip_declare,omitempty" yaml:"skip_declare,omitempty" toml:"skip_declare,omitempty"`

	prefix Identifier
}

// GoBuild is a placeholder for Go build tags.
type GoBuild struct {
	Tags string `json:"tags,omitempty" mapstructure:"tags,omitempty" yaml:"tags,omitempty" toml:"tags,omitempty"`
}

// GoPackage holds configuration details about the parent Go package that code will be generated for.
type GoPackage struct {
	Name string `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" validate:"required"`
	Path string `json:"path,omitempty" mapstructure:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}

// Prefix returns an Identifier for the enum's Go type alias in a normalized form.
func (g *GoType) Prefix() Identifier {
	if g.prefix.original == "" {
		g.prefix = NewIdentifier(g.Name)
	}

	return g.prefix
}

func emptyGo() *Go {
	return &Go{
		Type:    emptyType(),
		Build:   emptyBuild(),
		Package: emptyPackage(),
	}
}

func emptyType() *GoType {
	t := &GoType{}
	err := defaults.Set(t)
	if err != nil {
		panic(err)
	}
	return t
}

func emptyBuild() *GoBuild {
	t := &GoBuild{}
	err := defaults.Set(t)
	if err != nil {
		panic(err)
	}
	return t
}

func emptyPackage() *GoPackage {
	t := &GoPackage{}
	err := defaults.Set(t)
	if err != nil {
		panic(err)
	}
	return t
}
