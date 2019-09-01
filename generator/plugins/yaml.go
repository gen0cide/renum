package plugins

import "github.com/gen0cide/renum/generator/config"

// NewYAMLPlugin creates a new renum generator plugin to support the yaml.Marshaler and yaml.Unmarshaler interfaces.
func NewYAMLPlugin() Plugin {
	return &yamlPlugin{
		base: newBase("yaml", 22),
	}
}

type yamlPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *yamlPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Serializers.YAML
}

// Validate implements the Plugin interface.
func (p *yamlPlugin) Validate(c *config.Config) error {
	return nil
}
