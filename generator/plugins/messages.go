package plugins

import (
	"text/template"

	"github.com/gen0cide/renum/generator/config"
)

// NewMessagesPlugin creates a new renum generator plugin to implement the error interface on enum types.
func NewMessagesPlugin() Plugin {
	p := &messagesPlugin{
		base: newBase("messages", 14),
	}

	p.addFuncs(template.FuncMap{
		"mapify": p.mapify,
	})

	return p
}

type messagesPlugin struct {
	base
}

func (p *messagesPlugin) mapify(c *config.Config) (string, error) {
	return mapifyStringBuilder(c, func(e config.Element) interface{} {
		return e.Message
	})
}

// Enabled implements the Plugin interface.
func (p *messagesPlugin) Enabled(c *config.Config) bool {
	return c.Plugins.Error
}

// Validate implements the Plugin interface.
func (p *messagesPlugin) Validate(c *config.Config) error {
	return nil
}
