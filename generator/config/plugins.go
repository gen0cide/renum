package config

// PluginConfig holds the switches to define which templates will get rendered in the final code.
type PluginConfig struct {
	Codes       CodePlugins     `json:"codes,omitempty" mapstructure:"codes,omitempty" yaml:"codes,omitempty" toml:"codes,omitempty"`
	Error       bool            `json:"error,omitempty" mapstructure:"error,omitempty" yaml:"error,omitempty" toml:"error,omitempty"`
	Text        bool            `json:"text,omitempty" mapstructure:"text,omitempty" yaml:"text,omitempty" toml:"text,omitempty"`
	JSON        bool            `json:"json,omitempty" mapstructure:"json,omitempty" yaml:"json,omitempty" toml:"json,omitempty"`
	YAML        bool            `json:"yaml,omitempty" mapstructure:"yaml,omitempty" yaml:"yaml,omitempty" toml:"yaml,omitempty"`
	SQL         bool            `json:"sql,omitempty" mapstructure:"sql,omitempty" yaml:"sql,omitempty" toml:"sql,omitempty"`
	Flags       bool            `json:"flags,omitempty" mapstructure:"flags,omitempty" yaml:"flags,omitempty" toml:"flags,omitempty"`
	Pascal      bool            `json:"pascal,omitempty" mapstructure:"pascal,omitempty" yaml:"pascal,omitempty" toml:"pascal,omitempty"`
	Screaming   bool            `json:"screaming,omitempty" mapstructure:"screaming,omitempty" yaml:"screaming,omitempty" toml:"screaming,omitempty"`
	Description bool            `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	Namespace   NamespaceConfig `json:"namespace,omitempty" mapstructure:"namespace,omitempty" yaml:"namespace,omitempty" toml:"namespace,omitempty"`
	Docs        DocConfig       `json:"docs,omitempty" mapstructure:"docs,omitempty" yaml:"docs,omitempty" toml:"docs,omitempty"`
}

// NamespaceConfig is used to hold information relating to namespacing for the final template.
type NamespaceConfig struct {
	Enabled   bool   `json:"enabled,omitempty" mapstructure:"enabled,omitempty" yaml:"enabled,omitempty" toml:"enabled,omitempty"`
	Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty" yaml:"namespace,omitempty" toml:"namespace,omitempty"`
}
