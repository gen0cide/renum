package plugins

import (
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewSourcesPlugin creates a new renum generator plugin to implement the renum.Typer interface.
func NewSourcesPlugin() Plugin {
	p := &sourcesPlugin{
		base: newBase("sources", 13),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type sourcesPlugin struct {
	base
}

func (p *sourcesPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Go.PackagePath)
		buf.WriteString(dot)
		buf.WriteString(e.PrefixedPascal())
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *sourcesPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.SQL
}

// Validate implements the Plugin interface.
func (p *sourcesPlugin) Validate(c *config.Config) error {
	return nil
}
