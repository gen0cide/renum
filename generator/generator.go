//go:generate fileb0x assets.toml

package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/gen0cide/renum/generator/static"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
	"gopkg.in/go-playground/validator.v9"
)

const (
	skipHolder = `_`
)

var (
	_funcs = addFuncs(sprig.TxtFuncMap())
)

func addFuncs(fm template.FuncMap) template.FuncMap {
	fm["stringify_underscore"] = StringifyUnderscore
	fm["stringify_pascal"] = StringifyPascal
	fm["stringify_camel"] = StringifyCamel
	fm["stringify_screaming"] = StringifyScreaming
	fm["stringify_command"] = StringifyCommand
	fm["mapify_descriptions"] = MapifyDescriptions
	fm["mapify_messages"] = MapifyMessages
	fm["mapify_yarpc_codes"] = MapifyYARPCCodes
	fm["mapify_http_codes"] = MapifyHTTPCodes
	fm["mapify_os_exit_codes"] = MapifyOSExitCodes
	fm["mapify_full_kinds"] = MapifyFullKinds
	fm["mapify_full_paths"] = MapifyFullPaths
	fm["mapify_kinds"] = MapifyKinds
	fm["mapify_error_messages"] = MapifyErrorMessages
	fm["mapify"] = Mapify
	fm["mapify_pascal"] = MapifyPascal
	fm["mapify_camel"] = MapifyCamel
	fm["mapify_screaming"] = MapifyScreaming
	fm["mapify_command"] = MapifyCommand
	fm["unmapify"] = Unmapify
	fm["unmapify_pascal"] = UnmapifyPascal
	fm["unmapify_camel"] = UnmapifyCamel
	fm["unmapify_screaming"] = UnmapifyScreaming
	fm["unmapify_command"] = UnmapifyCommand
	fm["namify"] = Namify
	fm["valueify"] = Valueify
	return fm
}

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	t              *template.Template
	knownTemplates map[string]*template.Template
	Config         *Config
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator(config *Config) (*Generator, error) {
	if config == nil {
		return nil, errors.New("cannot create a generator with a nil config")
	}

	AddIdentifierInitialism(config.Go.Initialisms...)

	g := &Generator{
		knownTemplates: make(map[string]*template.Template),
		t:              template.New("generator"),
		Config:         config,
	}

	g.t.Funcs(_funcs)

	assets, err := static.WalkDirs("", false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to enumerate packed assets")
	}

	for _, assetName := range assets {
		if filepath.Ext(assetName) != ".tmpl" {
			continue
		}
		data, err := static.ReadFile(assetName)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read static asset %s", assetName)
		}

		g.t = template.Must(g.t.Parse(string(data)))
	}

	g.updateTemplates()

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

	g.Config.Go.Prefix = NewIdentifier(g.Config.Go.Name)

	unknownVal := Element{
		Ident:   NewIdentifier("undefined_enum_value"),
		Name:    "undefined_enum_value",
		Value:   0,
		Go:      g.Config.Go,
		Plugins: g.Config.Plugins,
	}

	unknownVal.Message = fmt.Sprintf("undefined enum value for type %s.%s", g.Config.Go.PackageName, g.Config.EnumID())
	unknownVal.Comment = fmt.Sprintf("%s is the default value for enum type %s. It is meant to be a placeholder and default for unknown values.", unknownVal.PrefixedPascal(), g.Config.EnumID())
	unknownVal.Description = fmt.Sprintf("This value is a default placeholder for any unknown type for the %s.%s enum.", g.Config.Go.PackageName, g.Config.EnumID())

	newValues := []Element{unknownVal}

	for idx, x := range g.Config.Values {
		x.Ident = NewIdentifier(x.Name)
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

	pkg := g.Config.Go.PackageName

	vBuff := bytes.NewBuffer([]byte{})
	err := g.t.ExecuteTemplate(vBuff, "header", map[string]interface{}{"package": pkg, "config": g.Config})
	if err != nil {
		return nil, errors.WithMessage(err, "Failed writing header")
	}

	err = g.t.ExecuteTemplate(vBuff, "definition", map[string]interface{}{"config": g.Config})
	if err != nil {
		return nil, errors.WithMessage(err, "Failed writing declaration")
	}

	errData := map[string]interface{}{
		"enum":      g.Config,
		"name":      g.Config.Go.Name,
		"lowercase": false,
		"names":     true,
		"config":    g.Config,
		"unknown":   !g.Config.Go.SkipUnknown,
	}

	err = g.t.ExecuteTemplate(vBuff, "enum", errData)
	if err != nil {
		return vBuff.Bytes(), errors.WithMessage(err, fmt.Sprintf("Failed writing enum data for enum: %q", g.Config.Go.Name))
	}

	formatted, err := imports.Process(pkg, vBuff.Bytes(), nil)
	if err != nil {
		err = fmt.Errorf("generate: error formatting code %s\n\n%s", err, vBuff.String())
	}
	return formatted, err
}

// updateTemplates will update the lookup map for validation checks that are
// allowed within the template engine.
func (g *Generator) updateTemplates() {
	for _, template := range g.t.Templates() {
		g.knownTemplates[template.Name()] = template
	}
}
