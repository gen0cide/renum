package plugins

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewErrorsPlugin creates a new renum generator plugin to implement the renum.Error interface on enum types.
func NewErrorsPlugin() Plugin {
	p := &errorsPlugin{
		base: newBase("errors", 18),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type errorsPlugin struct {
	base
}

func (p *errorsPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		buf := new(strings.Builder)
		fmt.Fprintf(buf, "%s.%s_%s (%d): %s", c.Namespace(), c.Go.Type.Prefix().Snake(), e.Snake(), e.Value, e.Message)
		return buf.String()
	})
}

// Enabled implements the Plugin interface.
func (p *errorsPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Renum.Error
}

// Validate implements the Plugin interface.
func (p *errorsPlugin) Validate(c *config.Config) error {
	return nil
}
