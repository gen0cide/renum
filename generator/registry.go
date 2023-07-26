package generator

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/gen0cide/renum/generator/config"
	"github.com/gen0cide/renum/generator/plugins"
	"github.com/gen0cide/renum/generator/static"
)

// Registry is a type that allows for Renderers to be bundled together
// and have operations performed in parallel on them.
type Registry struct {
	plugins map[string]plugins.Plugin
}

// NewEmptyRegistry creates a registry with an empty Plugin table.
func NewEmptyRegistry() *Registry {
	return &Registry{
		plugins: map[string]plugins.Plugin{},
	}
}

// NewDefaultRegistry returns a registry with access to all plugins within the renum library.
func NewDefaultRegistry() *Registry {
	return NewRegistryWithPlugins(plugins.All...)
}

// NewRegistryWithPlugins is used to create a pre-populated plugin registry by giving
// this constructor a list of plugins.Constructor types to create each plugin.
func NewRegistryWithPlugins(c ...plugins.Constructor) *Registry {
	reg := NewEmptyRegistry()
	for _, x := range c {
		p := x()

		reg.plugins[p.Name()] = p
	}

	return reg
}

// Add inserts the Plugin into it's table.
func (r *Registry) Add(rdr plugins.Plugin) {
	r.plugins[rdr.Name()] = rdr
}

// Delete removes a Plugin of a given name from the Registry's lookup table.
func (r *Registry) Delete(name string) {
	delete(r.plugins, name)
}

// Clone copies a registry, including it's current plugin table.
func (r *Registry) Clone() *Registry {
	newTable := make(map[string]plugins.Plugin, len(r.plugins))
	for k, v := range r.plugins {
		newTable[k] = v
	}

	return &Registry{
		plugins: newTable,
	}
}

// Build iterates all Plugin plugins to determine if they should be enabled, removing any that
// should not be enabled. For any remaining, it performs validation to ensure the Plugin plugin
// is compatible with the provided configuration, then called the Render() function to construct
// the code segment for that plugin.
func (r *Registry) Build(c *config.Config) error {
	if c == nil {
		return errors.New("registry cannot prune with a nil config")
	}

	renderers := make([]plugins.Plugin, len(r.plugins))
	idx := 0
	for _, v := range r.plugins {
		renderers[idx] = v
		idx++
	}

	for _, x := range renderers {
		if !x.Enabled(c) {
			delete(r.plugins, x.Name())
			continue
		}

		tmplData, err := static.ReadFile(fmt.Sprintf("%s.tmpl", x.Name()))
		if err != nil {
			return errors.Wrapf(err, "error reading template for plugin %s", x.Name())
		}

		err = x.Load(tmplData)
		if err != nil {
			return errors.Wrapf(err, "error loading template for plugin %s", x.Name())
		}

		err = x.Validate(c)
		if err != nil {
			return errors.Wrapf(err, "error validating plugin %s", x.Name())
		}

		err = x.Render(c)
		if err != nil {
			return errors.Wrapf(err, "error rendering plugin %s", x.Name())
		}
	}

	return nil
}

// Assemble is used to create the fully assembled Go source file containing the generated
// enum and all it's implementations.
func (r *Registry) Assemble(c *config.Config) ([]byte, error) {
	// List the enabled renderers and sort them by priority
	rlist := plugins.PluginList{}
	for _, x := range r.plugins {
		rlist = append(rlist, x)
	}
	sort.Sort(rlist)

	// enumerate the sorted list of renderers and append their data to the buffer
	buf := new(bytes.Buffer)
	for _, x := range rlist {
		data, err := x.Result()
		if err != nil {
			return nil, errors.Wrapf(err, "error exporting data for %s plugin", x.Name())
		}

		_, err = buf.Write(data)
		if err != nil {
			return nil, errors.Wrapf(err, "error appending buffer with data for %s plugin", x.Name())
		}

		buf.WriteString("\n\n")
	}

	buf.WriteString("/* \n --- BEGIN CONFIG DUMP ---\n\n")

	data, err := yaml.Marshal(c)
	if err != nil {
		return nil, errors.Wrapf(err, "error marshaling the configuration to YAML for embedding in the file")
	}

	_, err = buf.Write(data)
	if err != nil {
		return nil, errors.Wrapf(err, "could not embed the YAML configuration into the generated source")
	}

	buf.WriteString("\n\n\n --- END CONFIG DUMP ---\n\n*/\n")

	return buf.Bytes(), nil
}
