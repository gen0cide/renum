package plugins

import "github.com/gen0cide/renum/generator/config"

// NewStringerPlugin creates a new renum generator plugin to support the fmt.Stringer interface.
func NewStringerPlugin() Plugin {
	return &stringerPlugin{
		base: newBase("stringer", 6),
	}
}

type stringerPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *stringerPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *stringerPlugin) Validate(c *config.Config) error {
	return nil
}
