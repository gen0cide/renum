package plugins

import (
	"fmt"
	"strings"

	"github.com/gen0cide/renum/generator/config"
)

// Case is used to automate the creation of case specific values during rendering.
type Case int

const (
	// CaseSnake represents snake_case
	CaseSnake Case = iota

	// CasePascal represents PascalCase
	CasePascal

	// CaseScreaming represents SCREAMING_CASE
	CaseScreaming

	// CaseCamel represents camelCase
	CaseCamel

	// CaseCommand represents command-case
	CaseCommand

	// CaseDefault is an alias to the default, snake_case
	CaseDefault = CaseSnake
)

// String returns the case as a string representation.
func (c Case) String() string {
	switch c {
	case CaseSnake:
		return "Snake"
	case CaseScreaming:
		return "Screaming"
	case CaseCamel:
		return "Camel"
	case CasePascal:
		return "Pascal"
	case CaseCommand:
		return "Command"
	default:
		return CaseDefault.String()
	}
}

// Val returns an enum element's identifier that matches the appropriate case.
func (c Case) Val(e config.Element) string {
	switch c {
	case CaseSnake:
		return e.Snake()
	case CaseScreaming:
		return e.Screaming()
	case CaseCamel:
		return e.Camel()
	case CasePascal:
		return e.Pascal()
	case CaseCommand:
		return e.Command()
	default:
		return CaseDefault.String()
	}
}

func unmapifyForCase(c *config.Config, pref Case) string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "map[string]%s{\n", c.Go.Prefix.Pascal())
	label := stringifyIdentifier(c, pref)
	index := 0
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		nextIndex := index + len(pref.Val(val))
		fmt.Fprintf(buf, "\t%s[%d:%d]: %d,\n", label, index, nextIndex, val.Value)
		index = nextIndex
	}
	buf.WriteString(closeBrace)
	return buf.String()
}

func mapifyForCase(c *config.Config, pref Case) string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "map[%s]string{\n", c.Go.Prefix.Pascal())
	label := stringifyIdentifier(c, pref)
	index := 0
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		nextIndex := index + len(pref.Val(val))
		fmt.Fprintf(buf, "\t%d: %s[%d:%d],\n", val.Value, label, index, nextIndex)
		index = nextIndex
	}
	buf.WriteString(closeBrace)
	return buf.String()
}

func stringifyForCase(c *config.Config, pref Case) string {
	buf := new(strings.Builder)
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		buf.WriteString(pref.Val(val))
	}

	return buf.String()
}

func stringifyIdentifier(c *config.Config, pref Case) string {
	buf := new(strings.Builder)
	buf.WriteString("_")
	buf.WriteString(c.Go.Prefix.Pascal())
	buf.WriteString(pref.String())
	buf.WriteString("Name")
	return buf.String()
}
