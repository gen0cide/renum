//go:generate fileb0x assets.toml

package generator

import (
	"fmt"
	"strings"

	"github.com/gen0cide/renum/generator/config"

	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
	"gopkg.in/go-playground/validator.v9"
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

	config.AddIdentifierInitialism(cfg.Go.Initialisms...)

	g := &Generator{
		reg:    NewDefaultRegistry(),
		Config: cfg,
	}

	return g, nil
}

// Initialize builds an index of the supplied values, calculating things like names and offsets.
func (g *Generator) Initialize() error {
	validate := validator.New()
	err := validate.Struct(g.Config)
	if err != nil {
		return err
	}
	g.Config.Plugins.Codes.Simple = true

	if g.Config.Go.Name == "" {
		return errors.New("must provide a name for the go type to be used (e.g. FooType)")
	}

	if g.Config.Go.PackageName == "" {
		return errors.New("must provide a go package name for the output file")
	}

	err = g.Config.Go.Validate()
	if err != nil {
		return err
	}

	patchedNamespace := strings.Replace(g.Config.Go.PackagePath, `/`, `.`, -1)

	// TODO: validate that patchedNamespace actually conforms to /[a-zA-Z\-\_\.]/ semantics.

	g.Config.Plugins.Namespace.Namespace = patchedNamespace
	g.Config.Plugins.Namespace.Enabled = true

	if g.Config.Plugins.Docs.Enabled && g.Config.Go.PackagePath == "" {
		return errors.New("must provide go.package_path value if the docs plugin is enabled")
	}

	g.Config.Go.Prefix = config.NewIdentifier(g.Config.Go.Name)

	unknownVal := config.Element{
		Ident:   config.NewIdentifier("undefined_enum_value"),
		Name:    "undefined_enum_value",
		Value:   0,
		Go:      g.Config.Go,
		Plugins: g.Config.Plugins,
	}

	unknownVal.Message = fmt.Sprintf("undefined enum value for type %s.%s", g.Config.Go.PackageName, g.Config.EnumID())
	unknownVal.Comment = fmt.Sprintf("%s is the default value for enum type %s. It is meant to be a placeholder and default for unknown values.", unknownVal.PrefixedPascal(), g.Config.EnumID())
	unknownVal.Description = fmt.Sprintf("This value is a default placeholder for any unknown type for the %s.%s enum.", g.Config.Go.PackageName, g.Config.EnumID())

	newValues := []config.Element{unknownVal}

	for idx, x := range g.Config.Values {
		x.Ident = config.NewIdentifier(x.Name)
		x.Value = idx + 1
		x.Go = g.Config.Go
		x.Plugins = g.Config.Plugins
		newValues = append(newValues, x)
	}

	g.Config.Values = newValues

	return nil
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

	formatted, err := imports.Process(g.Config.Go.PackageName, code, nil)
	if err != nil {
		err = fmt.Errorf("generate: error formatting code %s\n\n%s", err, string(code))
	}

	return formatted, err
}
