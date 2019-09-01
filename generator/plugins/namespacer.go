package plugins

import (
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewNamespacerPlugin creates a new renum generator plugin to implement the renum.Namespacer interface.
func NewNamespacerPlugin() Plugin {
	p := &namespacerPlugin{
		base: newBase("namespacer", 19),
	}

	p.addFuncs(template.FuncMap{
		"mapify_paths": p.mapifyPaths,
		"mapify_ids":   p.mapifyIDs,
	})

	return p
}

type namespacerPlugin struct {
	base
}

func (p *namespacerPlugin) mapifyPaths(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Namespace())
		buf.WriteString(dot)
		buf.WriteString(c.Go.Type.Prefix().Snake())
		buf.WriteString(under)
		buf.WriteString(e.Snake())
		return buf.String()
	})
}

func (p *namespacerPlugin) mapifyIDs(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Go.Type.Prefix().Snake())
		buf.WriteString(under)
		buf.WriteString(e.Snake())
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *namespacerPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Renum.Namespacer
}

// Validate implements the Plugin interface.
func (p *namespacerPlugin) Validate(c *config.Config) error {
	return nil
}
