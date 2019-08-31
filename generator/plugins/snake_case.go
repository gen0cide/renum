package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewSnakeCasePlugin creates a new renum generator plugin to support snake_case enum strings.
func NewSnakeCasePlugin() Plugin {
	p := &snakeCasePlugin{
		base: newBase("snake_case", 7),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type snakeCasePlugin struct {
	base
}

func (p *snakeCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseSnake)
}

func (p *snakeCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseSnake)
}

func (p *snakeCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseSnake)
}

// Enabled implements the Plugin interface.
func (p *snakeCasePlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *snakeCasePlugin) Validate(c *config.Config) error {
	return nil
}
