package plugins

import (
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewSourcerPlugin creates a new renum generator plugin to implement the renum.Sourcer interface.
func NewSourcerPlugin() Plugin {
	p := &sourcerPlugin{
		base: newBase("sourcer", 13),
	}

	p.addFuncs(template.FuncMap{
		"mapify_export_refs":  p.mapifyExportRefs,
		"mapify_export_types": p.mapifyExportTypes,
	})

	return p

}

type sourcerPlugin struct {
	base
}

func (p *sourcerPlugin) mapifyExportRefs(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Go.Package.Path)
		buf.WriteString(dot)
		buf.WriteString(e.PrefixedPascal())
		return buf.String()
	})
}

func (p *sourcerPlugin) mapifyExportTypes(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Go.Package.Name)
		buf.WriteString(dot)
		buf.WriteString(e.PrefixedPascal())
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *sourcerPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Renum.Sourcer
}

// Validate implements the Plugin interface.
func (p *sourcerPlugin) Validate(c *config.Config) error {
	return nil
}
