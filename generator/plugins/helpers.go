package plugins

import (
	"fmt"
	"io"
	"strings"

	"github.com/gen0cide/renum/generator/config"
)

// ValueType is an alias to help automate the construction of code generated maps
// for enum lookups.
type ValueType int

const (
	// ValueTypeString is the default value type and is used to construct map[%s]string{} types.
	ValueTypeString ValueType = iota

	// ValueTypeInt is used to construct map[%s]int{} types.
	ValueTypeInt

	// ValueTypeYARPCError is used to construct map[%s]yarpcerrors.Code types.
	ValueTypeYARPCError
)

// String implements the fmt.Stringer interface.
func (v ValueType) String() string {
	switch v {
	case ValueTypeString:
		return "string"
	case ValueTypeInt:
		return "int"
	case ValueTypeYARPCError:
		return "yarpcerrors.Code"
	default:
		return ValueTypeString.String()
	}
}

// FormatValue is used to return the format string appropriate for the type of map being built.
func (v ValueType) FormatValue(buf io.Writer, values ...interface{}) (int, error) {
	switch v {
	case ValueTypeString:
		return fmt.Fprintf(buf, "\t%d: `%s`,\n", values...)
	case ValueTypeInt:
		return fmt.Fprintf(buf, "\t%d: %d,\n", values...)
	case ValueTypeYARPCError:
		return fmt.Fprintf(buf, "\t%d: %s,\n", values...)
	default:
		return ValueTypeString.FormatValue(buf, values...)
	}
}

// GetterFunc is a type alias to allow retrieval of element fields on the fly.
type GetterFunc func(e config.Element) interface{}

func mapifyStringBuilder(c *config.Config, getter GetterFunc) (string, error) {
	return mapBuilder(c, ValueTypeString, getter)
}

func mapBuilder(c *config.Config, vtype ValueType, getter GetterFunc) (string, error) {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "map[%s]%s{\n", c.Go.Type.Prefix().Pascal(), vtype.String())
	for _, val := range c.Values {
		if val.Pascal() == skipHolder {
			continue
		}

		_, err := vtype.FormatValue(buf, val.Value, getter(*val))
		if err != nil {
			return "", fmt.Errorf("error attempting to write map value %s for type %s: %v", val.PrefixedPascal(), vtype.String(), err)
		}
	}
	buf.WriteString(closeBrace)
	return buf.String(), nil
}
