package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewOSExitCodesPlugin creates a new renum generator plugin to implement the renum.ProcessResponder interface.
func NewOSExitCodesPlugin() Plugin {
	p := &osExitCodesPlugin{
		base: newBase("os_exit_codes", 27),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type osExitCodesPlugin struct {
	base
}

func (p *osExitCodesPlugin) mapify(c *config.Config) (string, error) {
	return mapBuilder(c, ValueTypeInt, func(e config.Element) interface{} {
		return e.OSExitCode()
	})
}

// Enabled implements the Plugin interface.
func (p *osExitCodesPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Codes.OSExit
}

// Validate implements the Plugin interface.
func (p *osExitCodesPlugin) Validate(c *config.Config) error {
	return nil
}
