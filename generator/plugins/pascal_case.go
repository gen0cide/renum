package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewPascalCasePlugin creates a new renum generator plugin to support PascalCase enum strings.
func NewPascalCasePlugin() Plugin {
	p := &pascalCasePlugin{
		base: newBase("pascal_case", 11),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type pascalCasePlugin struct {
	base
}

func (p *pascalCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CasePascal)
}

func (p *pascalCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CasePascal)
}

func (p *pascalCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CasePascal)
}

// Enabled implements the Plugin interface.
func (p *pascalCasePlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *pascalCasePlugin) Validate(c *config.Config) error {
	return nil
}
