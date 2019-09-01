package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewTrainCasePlugin creates a new renum generator plugin to support TRAIN-CASE enum strings.
func NewTrainCasePlugin() Plugin {
	p := &trainCasePlugin{
		base: newBase("train_case", 12),
	}

	p.addFuncs(template.FuncMap{
		"stringify": p.stringify,
		"mapify":    p.mapify,
		"unmapify":  p.unmapify,
	})

	return p
}

type trainCasePlugin struct {
	base
}

func (p *trainCasePlugin) stringify(c *config.Config) string {
	return stringifyForCase(c, CaseTrain)
}

func (p *trainCasePlugin) mapify(c *config.Config) string {
	return mapifyForCase(c, CaseTrain)
}

func (p *trainCasePlugin) unmapify(c *config.Config) string {
	return unmapifyForCase(c, CaseTrain)
}

// Enabled implements the Plugin interface.
func (p *trainCasePlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Cases.Train
}

// Validate implements the Plugin interface.
func (p *trainCasePlugin) Validate(c *config.Config) error {
	return nil
}
