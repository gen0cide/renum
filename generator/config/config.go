package config

import (
	"github.com/creasty/defaults"
	"github.com/pkg/errors"

	"github.com/gen0cide/renum"
)

// Config holds the top level configuration object used during generation.
type Config struct {
	OutputDir string       `json:"output_dir,omitempty" mapstructure:"output_dir,omitempty" yaml:"output_dir,omitempty" toml:"output_dir,omitempty" validate:"required,dir"`
	Go        GoConfig     `json:"go,omitempty" mapstructure:"go,omitempty" yaml:"go,omitempty" toml:"go,omitempty"`
	Plugins   PluginConfig `json:"plugins,omitempty" mapstructure:"plugins,omitempty" yaml:"plugins,omitempty" toml:"plugins,omitempty"`
	Values    []Element    `json:"values,omitempty" mapstructure:"values,omitempty" yaml:"values,omitempty" toml:"values,omitempty" validate:"gt=0"`
}

// NewConfig creates a new configuration, setting it's default values.
func NewConfig() (Config, error) {
	config := Config{}
	err := defaults.Set(&config)
	if err != nil {
		return config, errors.Wrap(err, "configuration failed to set defaults")
	}
	return config, nil
}

// EnumID is a template helper function that returns the enum's generated type.
func (c Config) EnumID() string {
	return c.Go.Prefix.Pascal()
}

// Version is a helper function for use during code generation in order to embed the
// renum library / command line generator version number into the generated source file.
func (c Config) Version() string {
	return renum.VersionString()
}
