package generator

// CodePlugins hold the configuration around which templates should be run.
type CodePlugins struct {
	Simple   bool         `json:"simple,omitempty" yaml:"simple,omitempty" toml:"simple,omitempty"`
	YARPC    bool         `json:"yarpc,omitempty" yaml:"yarpc,omitempty" toml:"yarpc,omitempty"`
	HTTP     bool         `json:"http,omitempty" yaml:"http,omitempty" toml:"http,omitempty"`
	OSExit   bool         `json:"os_exit,omitempty" yaml:"os_exit,omitempty" toml:"os_exit,omitempty"`
	Defaults CodeDefaults `json:"defaults,omitempty" yaml:"defaults,omitempty" toml:"defaults,omitempty"`
}

// CodeDefaults are used to hold preferences about what the values of coded responders should be.
type CodeDefaults struct {
	YARPC  *int `json:"yarpc,omitempty" yaml:"yarpc,omitempty" toml:"yarpc,omitempty"`
	HTTP   *int `json:"http,omitempty" yaml:"http,omitempty" toml:"http,omitempty"`
	OSExit *int `json:"os_exit,omitempty" yaml:"os_exit,omitempty" toml:"os_exit,omitempty"`
}
