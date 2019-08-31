package plugins

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewValuesPlugin creates a new renum generator plugin to add a slice of possible enum values into the package.
func NewValuesPlugin() Plugin {
	p := &valuesPlugin{
		base: newBase("values", 15),
	}

	p.addFuncs(template.FuncMap{
		"valueify": p.valueify,
	})

	return p
}

type valuesPlugin struct {
	base
}

func (p *valuesPlugin) valueify(c *config.Config) string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "[]%s{\n", c.EnumID())
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		fmt.Fprintf(buf, "\t%s,\n", val.PrefixedPascal())
	}
	buf.WriteString(closeBrace)
	return buf.String()
}

// Enabled implements the Plugin interface.
func (p *valuesPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *valuesPlugin) Validate(c *config.Config) error {
	return nil
}
