package plugins

import "github.com/gen0cide/renum/generator/config"

// NewDefinitionPlugin creates a new renum generator plugin to render the type alias for the renum.Enum base type.
func NewDefinitionPlugin() Plugin {
	return &definitionPlugin{
		base: newBase("definition", 2),
	}
}

type definitionPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *definitionPlugin) Enabled(c *config.Config) bool {
	return !c.Go.Type.SkipDeclare
}

// Validate implements the Plugin interface.
func (p *definitionPlugin) Validate(c *config.Config) error {
	return nil
}
