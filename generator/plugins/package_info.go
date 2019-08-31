package plugins

import "github.com/gen0cide/renum/generator/config"

// NewPackageInfoPlugin creates a new renum generator plugin to implement the renum.Typer interface.
func NewPackageInfoPlugin() Plugin {
	return &packageInfoPlugin{
		base: newBase("package_info", 6),
	}
}

type packageInfoPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *packageInfoPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *packageInfoPlugin) Validate(c *config.Config) error {
	return nil
}
