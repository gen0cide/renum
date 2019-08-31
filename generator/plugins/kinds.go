package plugins

import (
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewKindsPlugin creates a new renum generator plugin to implement renum.Typer interface methods.
func NewKindsPlugin() Plugin {
	p := &kindsPlugin{
		base: newBase("kinds", 10),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type kindsPlugin struct {
	base
}

func (p *kindsPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		buf.WriteString(c.Go.PackageName)
		buf.WriteString(dot)
		buf.WriteString(e.PrefixedPascal())
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *kindsPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *kindsPlugin) Validate(c *config.Config) error {
	return nil
}
