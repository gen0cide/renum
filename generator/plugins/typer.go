package plugins

import "github.com/gen0cide/renum/generator/config"

// NewTyperPlugin creates a new renum generator plugin to support the renum.Typer interface.
func NewTyperPlugin() Plugin {
	return &typerPlugin{
		base: newBase("typer", 10),
	}
}

type typerPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *typerPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *typerPlugin) Validate(c *config.Config) error {
	return nil
}
