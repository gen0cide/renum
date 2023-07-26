package config

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

var (
	pascalCaseRE = regexp.MustCompile(`^[A-Z][a-z]+(?:[A-Z][a-z]+)*$`)
	snakeCaseRE  = regexp.MustCompile(`^[a-z]+(_[a-z]+)*$`)
	namespaceRE  = regexp.MustCompile(`^[a-z0-9\.\_]+$`)
)

// HandlePresets is used to set defaults before fields are read in from the config.
func HandlePresets(preset string, c *Config) error {
	switch preset {
	case "error":
		c.Plugins.Renum.Error = true
		c.enableFullEnum()
	case "enum":
		c.enableFullEnum()
	case "base":
		c.enableBaseEnum()
	default:
		c.enableFullEnum()
	}

	return nil
}

// Validate attempts to validate the given config for errors
//
//nolint:gocyclo
func Validate(c *Config) error {
	if len(c.Initialisms) > 0 {
		AddIdentifierInitialism(c.Initialisms...)
	}

	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return errors.Wrap(err, "basic validation failed for the configuration")
	}

	// check if the output directory exists
	outdir, err := os.Stat(c.Output.Dir)
	if err != nil {
		return errors.Wrap(err, "output directory does not appear to be valid")
	}

	if !outdir.IsDir() {
		return errors.New("output directory is not a directory")
	}

	// check if name is valid pascal case
	if !pascalCaseRE.MatchString(c.Go.Type.Name) {
		return fmt.Errorf("go.type.name value %s does not appear to be valid PascalCase", c.Go.Type.Name)
	}

	// verify that we have everything we need for type declaration
	if !c.Go.Type.SkipDeclare {
		// do we have a comment?
		if c.Go.Type.Comment == "" {
			return errors.New("go.type.comment must have a valid string for type declaration")
		}

		// does the comment match the expectation for Go exported types?
		items := strings.Split(c.Go.Type.Comment, ` `)
		if items[0] != c.Go.Type.Name {
			return fmt.Errorf("go.type.comment is not a valid Go comment - must start with %s", c.Go.Type.Name)
		}
	}

	// ensure that we have a valid package name
	if c.Go.Package.Name == "" {
		return errors.New("go.package.name is a required field")
	}

	// ensure that if a package path was supplied that it matches the package name supplied
	if c.Go.Package.Path != "" && path.Base(c.Go.Package.Path) != c.Go.Package.Name {
		return errors.New("specified go.package.name does not match provided go.package.path basename")
	}

	// // check to see if we need to enable any settings due to a preset
	// switch c.Presets.Use {
	// case "error":
	// 	c.Plugins.Renum.Error = true
	// 	c.enableRenumEnum()
	// case "enum":
	// 	c.enableRenumEnum()
	// }

	// ensure we have package path set for required enum implementations
	if c.Plugins.Renum.Namespacer && c.Go.Package.Path == "" {
		return errors.New("must provide go.package.path to generate renum.Namespacer implementations")
	}

	if c.Plugins.Renum.Sourcer && c.Go.Package.Path == "" {
		return errors.New("must provide go.package.path to generate renum.Sourcer implementations")
	}

	if c.Plugins.Renum.Error && c.Go.Package.Path == "" {
		return errors.New("must provide go.package.path to generate renum.Error implementations")
	}

	// validate that the namespace we've been provided isn't garbage
	if c.Plugins.Renum.Namespacer {
		if len(c.Namespace()) < 3 {
			return errors.New("namespace was less than three characters in length - does your go.package.path contain invalid characters?")
		}

		if !namespaceRE.MatchString(c.Namespace()) {
			return fmt.Errorf("namespace generator could not produce a valid namespace - %s does not match /^[a-z0-9.-]+$/ regular expression", c.Namespace())
		}
	}

	// validate the go type name actually has a valid PascalCase representation
	c.Go.Type.prefix = NewIdentifier(c.Go.Type.Name)
	if c.Go.Type.Prefix().Pascal() != c.Go.Type.Name {
		return fmt.Errorf("go.type.name did not match expected PascalCase typing: provided %s, calculated %s", c.Go.Type.Name, c.Go.Type.Prefix().Pascal())
	}

	// create the initial enum value
	patientZero := &Element{
		Ident:   NewIdentifier("undefined_enum_value"),
		Name:    "undefined_enum_value",
		Value:   0,
		Message: fmt.Sprintf("undefined enum value for type %s.%s", c.Go.Package.Name, c.EnumID()),
		Config:  c,
	}

	patientZero.prefixed = NewIdentifier(fmt.Sprintf("%s_%s", c.Go.Type.Prefix().Snake(), patientZero.Ident.Snake()))
	patientZero.Description = fmt.Sprintf("%s is the default value for enum type %s. It is meant to be a place holder and default for unknown values.", patientZero.PrefixedPascal(), c.EnumID())

	enumValues := []*Element{patientZero}

	usedName := map[string]bool{}

	// now iterate the provided values and check conformity
	for idx, x := range c.Values {
		// ensure a name was provided
		if x.Name == "" {
			return fmt.Errorf("enum value %d did not provide a snake_case name", idx+1)
		}

		// ensure name was a valid snake_case identifier
		if !snakeCaseRE.MatchString(x.Name) {
			return fmt.Errorf("enum value %d (%s) does not have a valid snake_case name", idx+1, x.Name)
		}

		x.Ident = NewIdentifier(x.Name)

		// ensure the identity is unique
		if _, ok := usedName[x.Ident.Pascal()]; ok {
			return fmt.Errorf("enum value %d (%s) collides a pascal case identifier (%s) with another enum value", idx+1, x.Name, x.Ident.Pascal())
		}

		usedName[x.Ident.Pascal()] = true

		// calculate the prefix
		x.prefixed = NewIdentifier(fmt.Sprintf("%s_%s", c.Go.Type.Prefix().Snake(), x.Ident.Snake()))

		// set it's ID
		x.Value = idx + 1

		// ensure we have a description
		if c.Plugins.Renum.Descriptioner && x.Description == "" {
			if c.Opts.Strict {
				return fmt.Errorf("enum value %d (%s) did not provide a description, which is required to generate a renum.Descriptioner implementation", x.Value, x.Name)
			}

			if x.Comment != "" {
				x.Description = fmt.Sprintf("%s is an enum value for type %s. %s %s", x.PrefixedPascal(), c.EnumID(), x.PrefixedPascal(), x.Comment)
			}
		}

		if x.Description == "" && x.Comment == "" {
			return fmt.Errorf("enum value %d (%s) did not supply either a comment or a description field (at least one required)", x.Value, x.Name)
		}

		// ensure we have an error message if we need it
		if c.Plugins.Renum.Error && x.Message == "" {
			return fmt.Errorf("enum value %d (%s) did not provide a message, which is required to generate a renum.Error implementation", x.Value, x.Name)
		}

		// set the element's config reference and add it to the enum value list
		x.Config = c
		enumValues = append(enumValues, x)
	}

	// set the enumValues list and lets be on our way!
	c.Values = enumValues

	return nil
}
