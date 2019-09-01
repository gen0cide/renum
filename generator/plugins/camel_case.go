package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewCamelCasePlugin creates a new renum generator plugin to support camelCase enum strings.
func NewCamelCasePlugin() Plugin {
	p := &camelCasePlugin{
		base: newBase("camel_case", 12),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type camelCasePlugin struct {
	base
}

func (p *camelCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseCamel)
}

func (p *camelCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseCamel)
}

func (p *camelCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseCamel)
}

// Enabled implements the Plugin interface.
func (p *camelCasePlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Cases.Camel
}

// Validate implements the Plugin interface.
func (p *camelCasePlugin) Validate(c *config.Config) error {
	return nil
}
