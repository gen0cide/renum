//go:generate fileb0x assets.toml

package generator

import (
	"fmt"

	"github.com/gen0cide/renum/generator/config"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	Config *config.Config
	reg    *Registry
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator(cfg *config.Config) (*Generator, error) {
	if cfg == nil {
		return nil, errors.New("cannot create a generator with a nil config")
	}

	g := &Generator{
		reg:    NewDefaultRegistry(),
		Config: cfg,
	}

	return g, nil
}

// Initialize builds an index of the supplied values, calculating things like names and offsets.
func (g *Generator) Initialize() error {
	return config.Validate(g.Config)
}

// GenerateEnums renders the Go code of the enums.
func (g *Generator) GenerateEnums() ([]byte, error) {
	err := g.reg.Build(g.Config)
	if err != nil {
		return nil, err
	}

	code, err := g.reg.Assemble()
	if err != nil {
		return nil, err
	}

	formatted, err := imports.Process(g.Config.Go.Package.Name, code, &imports.Options{FormatOnly: true, Comments: true})
	if err != nil {
		err = fmt.Errorf("generate: error formatting code %s\n\n%s", err, string(code))
	}

	return formatted, err
}
