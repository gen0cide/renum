package config

var (
	_builtinYARPCCode  = 2
	_builtinHTTPCode   = 520
	_builtinOSExitCode = 1
)

// Element represents a single, unique value within an enum that's being generated.
type Element struct {
	Name        string     `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty" validate:"required"`
	Message     string     `json:"message,omitempty" mapstructure:"message,omitempty" yaml:"message,omitempty" toml:"message,omitempty"`
	Comment     string     `json:"comment,omitempty" mapstructure:"comment,omitempty" yaml:"comment,omitempty" toml:"comment,omitempty"`
	Description string     `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	Codes       CodeValues `json:"codes,omitempty" mapstructure:"codes,omitempty" yaml:"codes,omitempty" toml:"codes,omitempty"`
	Ident       Identifier `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
	Value       int        `json:"-" yaml:"-" toml:"-" mapstructure:"-"`
	Config      *Config    `json:"-" yaml:"-" toml:"-" mapstructure:"-"`

	prefixed Identifier
}

// Prefixed returns an identifier for the element with the enum type prefix attached.
func (e Element) Prefixed() Identifier {
	return e.prefixed
	// return NewIdentifier(fmt.Sprintf("%s_%s", e.Config.Go.Type.Prefix().Snake(), e.Snake()))
}

// PrefixedSnake is a template function that returns the a qualified name with the enum's type as a prefix (snake_case).
func (e Element) PrefixedSnake() string {
	return e.prefixed.Snake()
}

// PrefixedPascal is a template function that returns the a qualified name with the enum's type as a prefix (PascalCase).
func (e Element) PrefixedPascal() string {
	return e.prefixed.Pascal()
}

// PrefixedScreaming is a template function that returns the a qualified name with the enum's type as a prefix (SCREAMING_CASE).
func (e Element) PrefixedScreaming() string {
	return e.prefixed.Screaming()
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

// Train returns the enum's name in TRAIN-CASE form.
func (e Element) Train() string {
	return e.Ident.Train()
}

// Dotted returns the enum's name in dotted.case form.
func (e Element) Dotted() string {
	return e.Ident.Dotted()
}

// YARPCCode is a template helper function that returns the enum element's YARPC error code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (2)
func (e Element) YARPCCode() int {
	ret := _builtinYARPCCode
	if e.Config.Plugins.Codes.YARPC && e.Config.Plugins.Codes.Defaults.YARPC != nil {
		ret = *e.Config.Plugins.Codes.Defaults.YARPC
	}
	if e.Config.Plugins.Codes.YARPC && e.Codes.YARPC != nil {
		ret = *e.Codes.YARPC
	}

	return ret
}

// HTTPCode is a template helper function that returns the enum element's HTTP Status code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (520)
func (e Element) HTTPCode() int {
	ret := _builtinHTTPCode
	if e.Config.Plugins.Codes.HTTP && e.Config.Plugins.Codes.Defaults.HTTP != nil {
		ret = *e.Config.Plugins.Codes.Defaults.HTTP
	}
	if e.Config.Plugins.Codes.HTTP && e.Codes.HTTP != nil {
		ret = *e.Codes.HTTP
	}

	return ret
}

// OSExitCode is a template helper function that returns the enum element's OS Exit code. That
// value is determined by order of priority: explicitly set in config, default set in config, and finally, the built in default. (1)
func (e Element) OSExitCode() int {
	ret := _builtinOSExitCode
	if e.Config.Plugins.Codes.OSExit && e.Config.Plugins.Codes.Defaults.OSExit != nil {
		ret = *e.Config.Plugins.Codes.Defaults.OSExit
	}
	if e.Config.Plugins.Codes.OSExit && e.Codes.OSExit != nil {
		ret = *e.Codes.OSExit
	}

	return ret
}
