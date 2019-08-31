package generator

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

	"github.com/pkg/errors"

	"github.com/gen0cide/renum/generator/config"
	"github.com/gen0cide/renum/generator/plugins"
	"github.com/gen0cide/renum/generator/static"
)

// Registry is a type that allows for Renderers to be bundled together
// and have operations performed in parallel on them.
type Registry struct {
	sync.RWMutex

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
	r.Lock()
	defer r.Unlock()

	r.plugins[rdr.Name()] = rdr
}

// Delete removes a Plugin of a given name from the Registry's lookup table.
func (r *Registry) Delete(name string) {
	r.Lock()
	defer r.Unlock()

	delete(r.plugins, name)
}

// Clone copies a registry, including it's current plugin table.
func (r *Registry) Clone() *Registry {
	r.Lock()
	defer r.Unlock()

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

	errchan := make(chan error, 1)
	delchan := make(chan string, 1)
	finchan := make(chan struct{}, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(renderers))

	go func() {
		wg.Wait()
		finchan <- struct{}{}
	}()

	for _, x := range renderers {
		go execute(x, c, wg, delchan, errchan)
	}

	for {
		select {
		case err := <-errchan:
			return err
		case del := <-delchan:
			r.Delete(del)
			continue
		case <-finchan:
			return nil
		}
	}
}

// Assemble is used to create the fully assembled Go source file containing the generated
// enum and all it's implementations.
func (r *Registry) Assemble() ([]byte, error) {
	r.Lock()
	defer r.Unlock()

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

	return buf.Bytes(), nil
}

func execute(r plugins.Plugin, c *config.Config, wg *sync.WaitGroup, delchan chan string, errchan chan error) {
	defer wg.Done()

	if r == nil {
		errchan <- errors.New("plugin was nil")
		return
	}

	if !r.Enabled(c) {
		delchan <- r.Name()
		return
	}

	tmplData, err := static.ReadFile(fmt.Sprintf("%s.tmpl", r.Name()))
	if err != nil {
		errchan <- errors.Wrapf(err, "error reading template for plugin %s", r.Name())
		return
	}

	err = r.Load(tmplData)
	if err != nil {
		errchan <- errors.Wrapf(err, "error loading template for plugin %s", r.Name())
		return
	}

	err = r.Validate(c)
	if err != nil {
		errchan <- errors.Wrapf(err, "error validating plugin %s", r.Name())
		return
	}

	err = r.Render(c)
	if err != nil {
		errchan <- errors.Wrapf(err, "error rendering plugin %s", r.Name())
	}
}
