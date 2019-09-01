package plugins

import "github.com/gen0cide/renum/generator/config"

// NewSQLPlugin creates a new renum generator plugin to support the sql.Scanner and driver.Valuer interfaces.
func NewSQLPlugin() Plugin {
	return &sqlPlugin{
		base: newBase("sql", 23),
	}
}

type sqlPlugin struct {
	base
}

// Enabled implements the Plugin interface.
func (p *sqlPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Serializers.SQL
}

// Validate implements the Plugin interface.
func (p *sqlPlugin) Validate(c *config.Config) error {
	return nil
}
