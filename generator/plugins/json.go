package plugins

import "github.com/gen0cide/renum/generator/config"

// NewJSONPlugin creates a new renum generator plugin to implement json.Marshaler and json.Unmarshaler interfaces.
func NewJSONPlugin() Plugin {
	return &jsonPlugin{
		base: newBase("json", 21),
	}
}

type jsonPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *jsonPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.JSON
}

// Validate implements the Plugin interface.
func (p *jsonPlugin) Validate(c *config.Config) error {
	return nil
}
