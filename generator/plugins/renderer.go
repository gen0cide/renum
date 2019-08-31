package plugins

import (
	"sort"

	"github.com/gen0cide/renum/generator/config"
)

// Plugin defines how a plugin is constructed within Renum, allowing for the codegen to
// be based on a set of Plugin components. These then get executed in parallel, and have their
// result combined together.
type Plugin interface {
	// Name returns the unique identifier of the plugin. This will be used
	// as the primary key in the renum generator renderer registry.
	Name() string

	// Load is called to provide the Plugin with it's Go template for codegen by the generator.
	Load(b []byte) error

	// Enabled is called to let a Plugin deterine if it should be used for this configuration.
	Enabled(c *config.Config) bool

	// Validate is used to ensure the configuration meets the requirements for this Plugin.
	Validate(c *config.Config) error

	// Render should construct the generated code, typically holding the result in a buffer.
	Render(c *config.Config) error

	// Result should return the buffer's data, or an error if an unexpected condition arises.
	Result() ([]byte, error)

	// Priority is used to determine the insertion order of the Result into the final generated code.
	Priority() int
}

// PluginList implements a slice for Plugins that can be sorted.
type PluginList []Plugin

// Len implements the sort.Sort interface.
func (r PluginList) Len() int {
	return len(r)
}

// Swap implements the sort.Sort interface.
func (r PluginList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Less implements the sort.Sort interface.
func (r PluginList) Less(i, j int) bool {
	// First, check to see if priorities can be ordered by priority
	if r[i].Priority() != r[j].Priority() {
		return r[i].Priority() < r[j].Priority()
	}

	// Priority was the same, now sort by name.
	names := []string{r[i].Name(), r[j].Name()}
	sort.Strings(names)
	return names[0] == r[i].Name()
}

// Constructor is a type alias to denote the signature for creating plugins.
type Constructor func() Plugin
