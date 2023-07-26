package config

// Codes holds the configuration options relating to what Responder implementations to generate.
type Codes struct {
	GRPC     bool       `json:"grpc,omitempty" mapstructure:"grpc,omitempty" yaml:"grpc,omitempty" toml:"grpc,omitempty"`
	YARPC    bool       `json:"yarpc,omitempty" mapstructure:"yarpc,omitempty" yaml:"yarpc,omitempty" toml:"yarpc,omitempty"`
	HTTP     bool       `json:"http,omitempty" mapstructure:"http,omitempty" yaml:"http,omitempty" toml:"http,omitempty"`
	Errno    bool       `json:"errno,omitempty" mapstructure:"errno,omitempty" yaml:"errno,omitempty" toml:"errno,omitempty"`
	OSExit   bool       `json:"os_exit,omitempty" mapstructure:"os_exit,omitempty" yaml:"os_exit,omitempty" toml:"os_exit,omitempty"`
	Defaults CodeValues `json:"defaults,omitempty" mapstructure:"defaults,omitempty" yaml:"defaults,omitempty" toml:"defaults,omitempty"`
}

// CodeValues holds user supplied values for Code Responders to use.
type CodeValues struct {
	GRPC   *int `json:"grpc,omitempty" mapstructure:"grpc,omitempty" yaml:"grpc,omitempty" toml:"grpc,omitempty"`
	YARPC  *int `json:"yarpc,omitempty" mapstructure:"yarpc,omitempty" yaml:"yarpc,omitempty" toml:"yarpc,omitempty"`
	HTTP   *int `json:"http,omitempty" mapstructure:"http,omitempty" yaml:"http,omitempty" toml:"http,omitempty"`
	Errno  *int `json:"errno,omitempty" mapstructure:"errno,omitempty" yaml:"errno,omitempty" toml:"errno,omitempty"`
	OSExit *int `json:"os_exit,omitempty" mapstructure:"os_exit,omitempty" yaml:"os_exit,omitempty" toml:"os_exit,omitempty"`
}
