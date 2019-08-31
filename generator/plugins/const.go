package plugins

import "github.com/gen0cide/renum/generator/config"

// NewConstPlugin creates a new renum generator plugin to create the const enum definitions.
func NewConstPlugin() Plugin {
	return &constPlugin{
		base: newBase("const", 3),
	}
}

type constPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *constPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *constPlugin) Validate(c *config.Config) error {
	return nil
}
