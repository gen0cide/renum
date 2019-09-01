package config

import (
	"fmt"
	"strings"

	"github.com/creasty/defaults"
	"github.com/gen0cide/renum"
)

// Config holds the primary configuration for the renum code generator.
type Config struct {
	Output      *Output     `json:"output,omitempty" mapstructure:"output,omitempty" yaml:"output,omitempty" toml:"output,omitempty"`
	Go          *Go         `json:"go,omitempty" mapstructure:"go,omitempty" yaml:"go,omitempty" toml:"go,omitempty"`
	Initialisms Initialisms `json:"initialisms,omitempty" mapstructure:"initialisms,omitempty" yaml:"initialisms,omitempty" toml:"initialisms,omitempty"`
	Presets     *Presets    `json:"presets,omitempty" mapstructure:"presets,omitempty" yaml:"presets,omitempty" toml:"presets,omitempty"`
	Plugins     *Plugins    `json:"plugins,omitempty" mapstructure:"plugins,omitempty" yaml:"plugins,omitempty" toml:"plugins,omitempty"`
	Values      []*Element  `json:"values,omitempty" mapstructure:"values,omitempty" yaml:"values,omitempty" toml:"values,omitempty" validate:"gt=0"`

	namespace string
}

// NewEmptyConfig creates a new empty renum configuration object.
func NewEmptyConfig() *Config {
	cfg := &Config{
		Output:      emptyOutput(),
		Go:          emptyGo(),
		Initialisms: Initialisms{},
		Presets:     emptyPresets(),
		Plugins:     emptyPlugins(),
		Values:      []*Element{},
	}

	err := defaults.Set(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

// Namespace returns the renum codegenerated package namespace.
func (c *Config) Namespace() string {
	if c.namespace == "" {
		c.namespace = strings.ToLower(strings.Replace(strings.Replace(c.Go.Package.Path, `/`, `.`, -1), `-`, `_`, -1))
	}

	return c.namespace
}

// EnumID returns the name of the enum's Go type alias in PascalCase.
func (c *Config) EnumID() string {
	return c.Go.Type.Prefix().Pascal()
}

// Version returns the Renum version string for this configuration.
func (c *Config) Version() string {
	return renum.VersionString()
}

// OutputFilename returns the name of the destination's filename to use.
func (c *Config) OutputFilename() string {
	if c.Output.Filename != "" {
		return c.Output.Filename
	}

	return fmt.Sprintf("generated_%s.go", c.Go.Type.Prefix().Ident().Underscore().Pluralize().String())
}

func (c *Config) enableRenumEnum() {
	c.Plugins.Renum.Namespacer = true
	c.Plugins.Renum.Sourcer = true
	c.Plugins.Renum.Typer = true
	c.Plugins.Renum.Coder = true
	c.Plugins.Renum.Descriptioner = true
	c.Plugins.Renum.Caser = true
	c.Plugins.Cases.Pascal = true
	c.Plugins.Cases.Camel = true
	c.Plugins.Cases.Screaming = true
	c.Plugins.Cases.Command = true
	c.Plugins.Cases.Dotted = true
	c.Plugins.Cases.Train = true
	c.Plugins.Serializers.Text = true
	c.Plugins.Serializers.JSON = true
	c.Plugins.Serializers.YAML = true
	c.Plugins.Serializers.CSV = true
	c.Plugins.Serializers.SQL = true
	c.Plugins.Serializers.Flags = true
	c.Plugins.Serializers.Gob = true
	c.Plugins.Serializers.Binary = true
}
