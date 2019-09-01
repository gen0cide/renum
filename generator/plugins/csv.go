package plugins

import "github.com/gen0cide/renum/generator/config"

// NewCSVPlugin creates a new renum generator plugin to implement csv.Marshaler and csv.Unmarshaler interfaces.
func NewCSVPlugin() Plugin {
	return &csvPlugin{
		base: newBase("csv", 22),
	}
}

type csvPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *csvPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Serializers.CSV
}

// Validate implements the Plugin interface.
func (p *csvPlugin) Validate(c *config.Config) error {
	return nil
}
