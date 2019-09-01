package plugins

import "github.com/gen0cide/renum/generator/config"

// NewTextPlugin creates a new renum generator plugin to support the encoding.TextMarshaler and encoding.TextUnmarshaler interfaces.
func NewTextPlugin() Plugin {
	return &textPlugin{
		base: newBase("text", 20),
	}
}

type textPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *textPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Serializers.Text
}

// Validate implements the Plugin interface.
func (p *textPlugin) Validate(c *config.Config) error {
	return nil
}
