package plugins

import "github.com/gen0cide/renum/generator/config"

// NewCoderPlugin creates a new renum generator plugin to support the renum.Coder interface.
func NewCoderPlugin() Plugin {
	return &coderPlugin{
		base: newBase("coder", 5),
	}
}

type coderPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *coderPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Renum.Coder
}

// Validate implements the Plugin interface.
func (p *coderPlugin) Validate(c *config.Config) error {
	return nil
}
