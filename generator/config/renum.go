package config

// Renum holds configuration options about what renum interface types to implement.
type Renum struct {
	Namespacer    bool `json:"namespacer,omitempty" mapstructure:"namespacer,omitempty" yaml:"namespacer,omitempty" toml:"namespacer,omitempty"`
	Sourcer       bool `json:"sourcer,omitempty" mapstructure:"sourcer,omitempty" yaml:"sourcer,omitempty" toml:"sourcer,omitempty"`
	Typer         bool `json:"typer,omitempty" mapstructure:"typer,omitempty" yaml:"typer,omitempty" toml:"typer,omitempty"`
	Coder         bool `json:"coder,omitempty" mapstructure:"coder,omitempty" yaml:"coder,omitempty" toml:"coder,omitempty"`
	Descriptioner bool `json:"descriptioner,omitempty" mapstructure:"descriptioner,omitempty" yaml:"descriptioner,omitempty" toml:"descriptioner,omitempty"`
	Caser         bool `json:"caser,omitempty" mapstructure:"caser,omitempty" yaml:"caser,omitempty" toml:"caser,omitempty"`
	Error         bool `json:"error,omitempty" mapstructure:"error,omitempty" yaml:"error,omitempty" toml:"error,omitempty"`
}
