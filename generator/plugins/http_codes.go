package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewHTTPCodesPlugin creates a new renum generator plugin to implement the renum.HTTPResponder interface.
func NewHTTPCodesPlugin() Plugin {
	p := &httpCodesPlugin{
		base: newBase("http_codes", 26),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type httpCodesPlugin struct {
	base
}

func (p *httpCodesPlugin) mapify(c *config.Config) (string, error) {
	return mapBuilder(c, ValueTypeInt, func(e config.Element) interface{} {
		return e.HTTPCode()
	})
}

// Enabled implements the Plugin interface.
func (p *httpCodesPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Codes.HTTP
}

// Validate implements the Plugin interface.
func (p *httpCodesPlugin) Validate(c *config.Config) error {
	return nil
}
