package plugins

import "github.com/gen0cide/renum/generator/config"

// NewFlagsPlugin creates a new renum generator plugin to support golang flag.Value interface.
func NewFlagsPlugin() Plugin {
	return &flagsPlugin{
		base: newBase("flags", 25),
	}
}

type flagsPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *flagsPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Serializers.Flags
}

// Validate implements the Plugin interface.
func (p *flagsPlugin) Validate(c *config.Config) error {
	return nil
}
