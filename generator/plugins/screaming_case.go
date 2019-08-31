package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewScreamingCasePlugin creates a new renum generator plugin to support SCREAMING_CASE enum strings.
func NewScreamingCasePlugin() Plugin {
	p := &screamingCasePlugin{
		base: newBase("screaming_case", 9),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type screamingCasePlugin struct {
	base
}

func (p *screamingCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseScreaming)
}

func (p *screamingCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseScreaming)
}

func (p *screamingCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseScreaming)
}

// Enabled implements the Plugin interface.
func (p *screamingCasePlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *screamingCasePlugin) Validate(c *config.Config) error {
	return nil
}
