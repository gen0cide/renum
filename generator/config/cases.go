package config

// Cases holds the configuration options to what casing generation should be enabled.
type Cases struct {
	Pascal    bool `json:"pascal,omitempty" mapstructure:"pascal,omitempty" yaml:"pascal,omitempty" toml:"pascal,omitempty"`
	Camel     bool `json:"camel,omitempty" mapstructure:"camel,omitempty" yaml:"camel,omitempty" toml:"camel,omitempty"`
	Screaming bool `json:"screaming,omitempty" mapstructure:"screaming,omitempty" yaml:"screaming,omitempty" toml:"screaming,omitempty"`
	Command   bool `json:"command,omitempty" mapstructure:"command,omitempty" yaml:"command,omitempty" toml:"command,omitempty"`
	Dotted    bool `json:"dotted,omitempty" mapstructure:"dotted,omitempty" yaml:"dotted,omitempty" toml:"dotted,omitempty"`
	Train     bool `json:"train,omitempty" mapstructure:"train,omitempty" yaml:"train,omitempty" toml:"train,omitempty"`
}
