package plugins

import "github.com/gen0cide/renum/generator/config"

// NewHeaderPlugin creates a new renum generator plugin to generate the package and import declaration on the generated code.
func NewHeaderPlugin() Plugin {
	return &headerPlugin{
		base: newBase("header", 1),
	}
}

type headerPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *headerPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *headerPlugin) Validate(c *config.Config) error {
	return nil
}
