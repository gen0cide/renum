package config

import "github.com/creasty/defaults"

// Presets holds a configuration stanza for allowing users to define a preset group of settings.
type Presets struct {
	Use string `json:"use,omitempty" mapstructure:"use,omitempty" yaml:"use,omitempty" toml:"use,omitempty" default:"enum" validate:"required,oneof=none enum error"`
}

func emptyPresets() *Presets {
	t := &Presets{}
	err := defaults.Set(t)
	if err != nil {
		panic(err)
	}
	return t
}
