package config

import (
	"fmt"
)

var (
	_builtinYARPCCode  = 2
	_builtinHTTPCode   = 520
	_builtinOSExitCode = 1
)

// Element represents a single, unique value within an enum that's being generated.
type Element struct {
	Name        string       `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Message     string       `json:"message,omitempty" mapstructure:"message,omitempty" yaml:"message,omitempty" toml:"message,omitempty"`
	Description string       `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	Comment     string       `json:"comment,omitempty" mapstructure:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty" validate:"required"`
	Codes       CodeDefaults `json:"codes,omitempty" mapstructure:"codes,omitempty" yaml:"codes,omitempty" toml:"codes,omitempty"`
	Docs        DocDef       `json:"docs,omitempty" mapstructure:"docs,omitempty" yaml:"docs,omitempty" toml:"docs,omitempty"`
	Ident       Identifier   `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
	Value       int          `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
	Go          GoConfig     `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
	Plugins     PluginConfig `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
}

// PrefixedSnake is a template function that returns the a qualified name with the enum's type as a prefix (snake_case).
func (e Element) PrefixedSnake() string {
	return fmt.Sprintf("%s_%s", e.Go.Prefix.Snake(), e.Snake())
}

// PrefixedPascal is a template function that returns the a qualified name with the enum's type as a prefix (PascalCase).
func (e Element) PrefixedPascal() string {
	return fmt.Sprintf("%s%s", e.Go.Prefix.Pascal(), e.Pascal())
}

// PrefixedScreaming is a template function that returns the a qualified name with the enum's type as a prefix (SCREAMING_CASE).
func (e Element) PrefixedScreaming() string {
	return fmt.Sprintf("%s_%s", e.Go.Prefix.Screaming(), e.Screaming())
}

// Camel returns the enum's name in camelCase form.
func (e Element) Camel() string {
	return e.Ident.Camel()
}

// Command returns the enums name in command-case form.
func (e Element) Command() string {
	return e.Ident.Command()
}

// Snake returns the enum's name in snake_case form.
func (e Element) Snake() string {
	return e.Ident.Snake()
}

// Pascal returns the enum's name in PascalCase form.
func (e Element) Pascal() string {
	return e.Ident.Pascal()
}

// Screaming returns the enum's name in SCREAMING_CASE form.
func (e Element) Screaming() string {
	return e.Ident.Screaming()
}

// YARPCCode is a template helper function that returns the enum element's YARPC error code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (2)
func (e Element) YARPCCode() int {
	ret := _builtinYARPCCode
	if e.Plugins.Codes.YARPC && e.Plugins.Codes.Defaults.YARPC != nil {
		ret = *e.Plugins.Codes.Defaults.YARPC
	}
	if e.Plugins.Codes.YARPC && e.Codes.YARPC != nil {
		ret = *e.Codes.YARPC
	}

	return ret
}

// HTTPCode is a template helper function that returns the enum element's HTTP Status code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (520)
func (e Element) HTTPCode() int {
	ret := _builtinHTTPCode
	if e.Plugins.Codes.HTTP && e.Plugins.Codes.Defaults.HTTP != nil {
		ret = *e.Plugins.Codes.Defaults.HTTP
	}
	if e.Plugins.Codes.HTTP && e.Codes.HTTP != nil {
		ret = *e.Codes.HTTP
	}

	return ret
}

// OSExitCode is a template helper function that returns the enum element's OS Exit code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (1)
func (e Element) OSExitCode() int {
	ret := _builtinOSExitCode
	if e.Plugins.Codes.OSExit && e.Plugins.Codes.Defaults.OSExit != nil {
		ret = *e.Plugins.Codes.Defaults.OSExit
	}
	if e.Plugins.Codes.OSExit && e.Codes.OSExit != nil {
		ret = *e.Codes.OSExit
	}

	return ret
}
