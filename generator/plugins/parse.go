package plugins

import "github.com/gen0cide/renum/generator/config"

// NewParsePlugin creates a new renum generator plugin to implement the required Parse{{Type}} function.
func NewParsePlugin() Plugin {
	return &parsePlugin{
		base: newBase("parse", 4),
	}
}

type parsePlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *parsePlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *parsePlugin) Validate(c *config.Config) error {
	return nil
}
