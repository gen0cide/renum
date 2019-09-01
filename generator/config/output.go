package config

// Output holds the configuration details about where to write the generated code.
type Output struct {
	Dir       string `json:"dir,omitempty" mapstructure:"dir,omitempty" yaml:"dir,omitempty" toml:"dir,omitempty" validate:"required,dir"`
	Filename  string `json:"filename,omitempty" mapstructure:"filename,omitempty" yaml:"filename,omitempty" toml:"filename,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty" mapstructure:"overwrite,omitempty" yaml:"overwrite,omitempty" toml:"overwrite,omitempty"`
}

func emptyOutput() *Output {
	return &Output{}
}
