package plugins

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewNamesPlugin creates a new renum generator plugin to create a string slice of enum values.
func NewNamesPlugin() Plugin {
	p := &namesPlugin{
		base: newBase("names", 16),
	}

	p.addFuncs(template.FuncMap{
		"namify": p.namify,
	})

	return p
}

type namesPlugin struct {
	base
}

func (p *namesPlugin) namify(c *config.Config) string {
	label := stringifyIdentifier(c, CaseSnake)
	buf := new(strings.Builder)
	buf.WriteString("[]string{\n")
	index := 0
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		nextIndex := index + len(val.Snake())
		fmt.Fprintf(buf, "\t%s[%d:%d],\n", label, index, nextIndex)
		index = nextIndex
	}
	buf.WriteString(closeBrace)
	return buf.String()
}

// Enabled implements the Plugin interface.
func (p *namesPlugin) Enabled(c *config.Config) bool {
	return true
}

// Validate implements the Plugin interface.
func (p *namesPlugin) Validate(c *config.Config) error {
	return nil
}
