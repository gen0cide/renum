package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewDescriptionsPlugin is used to create a renum generator plugin for implementing the renum.Descriptioner interface on enum types.
func NewDescriptionsPlugin() Plugin {
	p := &descriptionsPlugin{
		base: newBase("descriptions", 17),
	}
	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type descriptionsPlugin struct {
	base
}

func (p *descriptionsPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		return e.Description
	})
}

// Enabled implements the Plugin interface.
func (p *descriptionsPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Description
}

// Validate implements the Plugin interface.
func (p *descriptionsPlugin) Validate(c *config.Config) error {
	return nil
}
