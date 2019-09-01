package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewDescriptionerPlugin is used to create a renum generator plugin for implementing the renum.Descriptioner interface on enum types.
func NewDescriptionerPlugin() Plugin {
	p := &descriptionerPlugin{
		base: newBase("descriptioner", 17),
	}
	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type descriptionerPlugin struct {
	base
}

func (p *descriptionerPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		return e.Description
	})
}

// Enabled implements the Plugin interface.
func (p *descriptionerPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Renum.Descriptioner
}

// Validate implements the Plugin interface.
func (p *descriptionerPlugin) Validate(c *config.Config) error {
	return nil
}
