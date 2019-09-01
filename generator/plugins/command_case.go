package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewCommandCasePlugin creates a new renum generator plugin to support command-case enum strings.
func NewCommandCasePlugin() Plugin {
	p := &commandCasePlugin{
		base: newBase("command_case", 8),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type commandCasePlugin struct {
	base
}

func (p *commandCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseCommand)
}

func (p *commandCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseCommand)
}

func (p *commandCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseCommand)
}

// Enabled implements the Plugin interface.
func (p *commandCasePlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Cases.Command
}

// Validate implements the Plugin interface.
func (p *commandCasePlugin) Validate(c *config.Config) error {
	return nil
}
