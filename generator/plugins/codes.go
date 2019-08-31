package plugins

import "github.com/gen0cide/renum/generator/config"

// NewCodesPlugin creates a new renum generator plugin to support the renum.Coder interface.
func NewCodesPlugin() Plugin {
	return &codesPlugin{
		base: newBase("codes", 5),
	}
}

type codesPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *codesPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *codesPlugin) Validate(c *config.Config) error {
	return nil
}
