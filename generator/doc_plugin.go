package generator

// DocConfig holds information about the docs automatic generation plugin. (unused currently, work in progress).
type DocConfig struct {
	Enabled    bool        `json:"enabled,omitempty" mapstructure:"enabled,omitempty" yaml:"enabled,omitempty" toml:"enabled,omitempty"`
	References bool        `json:"references,omitempty" mapstructure:"references,omitempty" yaml:"references,omitempty" toml:"references,omitempty"`
	OutputDir  string      `json:"output_dir,omitempty" mapstructure:"output_dir,omitempty" yaml:"output_dir,omitempty" toml:"output_dir,omitempty" default:"docs"`
	Godoc      GodocConfig `json:"godoc,omitempty" mapstructure:"godoc,omitempty" yaml:"godoc,omitempty" toml:"godoc,omitempty"`
}

// GodocConfig holds information about the relavent Godoc paths to use when generating documentation.
type GodocConfig struct {
	Enabled bool   `json:"enabled,omitempty" mapstructure:"enabled,omitempty" yaml:"enabled,omitempty" toml:"enabled,omitempty"`
	BaseURL string `json:"base_url,omitempty" mapstructure:"base_url,omitempty" yaml:"base_url,omitempty" toml:"base_url,omitempty" default:"https://godoc.org/" validate:"required,url"`
}

// DocDef allows for enums to reference existing code that will be embedded in the documentation for the enum.
type DocDef struct {
	RefType string `json:"ref_type,omitempty" mapstructure:"ref_type,omitempty" yaml:"ref_type,omitempty" toml:"ref_type,omitempty"`
	RefFunc string `json:"ref_func,omitempty" mapstructure:"ref_func,omitempty" yaml:"ref_func,omitempty" toml:"ref_func,omitempty"`
}
