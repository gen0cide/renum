package config

// Plugins is a container structure that holds all the configuration settings for what codegen
// plugins should be used by the renum generator.
type Plugins struct {
	Renum       Renum       `json:"renum,omitempty" mapstructure:"renum,omitempty" yaml:"renum,omitempty" toml:"renum,omitempty"`
	Cases       Cases       `json:"cases,omitempty" mapstructure:"cases,omitempty" yaml:"cases,omitempty" toml:"cases,omitempty"`
	Serializers Serializers `json:"serializers,omitempty" mapstructure:"serializers,omitempty" yaml:"serializers,omitempty" toml:"serializers,omitempty"`
	Codes       Codes       `json:"codes,omitempty" mapstructure:"codes,omitempty" yaml:"codes,omitempty" toml:"codes,omitempty"`
}
