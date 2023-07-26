package config

// Options is a configuration container that is used to hold some various flags surrounding
// the execution of renum.
type Options struct {
	Strict bool `json:"strict,omitempty" mapstructure:"strict,omitempty" yaml:"strict,omitempty" toml:"strict,omitempty" default:"true"`
}
