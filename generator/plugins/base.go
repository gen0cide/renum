package plugins

import (
	"bytes"
	"errors"
	"sync"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/gen0cide/renum/generator/config"
)

const (
	skipHolder = `_`
	closeBrace = `}`
	dot        = `.`
	under      = `_`
)

var (
	// ErrNilTemplate is thrown when a renderer plugin attempts an operation that cannot be performed
	// because the renderer's template has not been loaded.
	ErrNilTemplate = errors.New("renderer plugin has nil template")

	// ErrEmptyResult is thrown when a renderer plugin attempts to export the Result and finds no
	// data in it's buffer.
	ErrEmptyResult = errors.New("renderer plugin returned an empty result")
)

type base struct {
	mutex    *sync.RWMutex
	t        *template.Template
	priority int
	name     string
	buf      *bytes.Buffer
	funcs    template.FuncMap
}

// newBase creates a new base plugin.
func newBase(name string, priority int) base {
	b := base{
		name:     name,
		priority: priority,
		buf:      new(bytes.Buffer),
		funcs:    sprig.TxtFuncMap(),
		mutex:    new(sync.RWMutex),
	}

	return b
}

// addFuncs is used to append functions to the template function map.
func (b *base) addFuncs(tm template.FuncMap) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.funcs == nil {
		b.funcs = template.FuncMap{}
	}
	for k, v := range tm {
		b.funcs[k] = v
	}
}

// Name implements the Plugin interface.
func (b *base) Name() string {
	return b.name
}

// Priority implements the Plugin interface.
func (b *base) Priority() int {
	if b.priority == 0 {
		return 100
	}

	return b.priority
}

// Load implements the Plugin interface.
func (b *base) Load(data []byte) error {
	t, err := template.New(b.name).Funcs(b.funcs).Parse(string(data))
	if err != nil {
		return err
	}

	b.t = t
	return nil
}

// Render implements the Plugin interface.
func (b *base) Render(c *config.Config) error {
	if b.t == nil {
		return ErrNilTemplate
	}

	if b.buf == nil {
		b.buf = new(bytes.Buffer)
	}

	err := b.t.Execute(b.buf, c)
	if err != nil {
		return err
	}

	return nil
}

// Result implements the Plugin interface.
func (b *base) Result() ([]byte, error) {
	if b.buf.Len() == 0 {
		return nil, ErrEmptyResult
	}

	return b.buf.Bytes(), nil
}
