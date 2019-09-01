package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewDottedCasePlugin creates a new renum generator plugin to support dotted.case enum strings.
func NewDottedCasePlugin() Plugin {
	p := &dottedCasePlugin{
		base: newBase("dotted_case", 12),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type dottedCasePlugin struct {
	base
}

func (p *dottedCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseDotted)
}

func (p *dottedCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseDotted)
}

func (p *dottedCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseDotted)
}

// Enabled implements the Plugin interface.
func (p *dottedCasePlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Cases.Dotted
}

// Validate implements the Plugin interface.
func (p *dottedCasePlugin) Validate(c *config.Config) error {
	return nil
}
