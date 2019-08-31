package plugins

import (
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewPathsPlugin creates a new renum generator plugin to implement the renum.Typer interface.
func NewPathsPlugin() Plugin {
	p := &pathsPlugin{
		base: newBase("paths", 19),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type pathsPlugin struct {
	base
}

func (p *pathsPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Plugins.Namespace.Namespace)
		buf.WriteString(dot)
		buf.WriteString(c.Go.Prefix.Snake())
		buf.WriteString(under)
		buf.WriteString(e.Snake())
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *pathsPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *pathsPlugin) Validate(c *config.Config) error {
	return nil
}
