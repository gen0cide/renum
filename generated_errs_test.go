// Code generated by renum (github.com/gen0cide/renum)
// DO NOT EDIT!
// renum v0.0.5-b63b6da

package renum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"go.uber.org/yarpc/yarpcerrors"
)

// ErrTest is a generated type alias for the ErrTest enum.
type ErrTest int

const (
	// ErrTestUndefinedEnumValue is an enum value for type ErrTest.
	// ErrTestUndefinedEnumValue is the default value for enum type ErrTest. It is meant to be a placeholder and default for unknown values.
	// This value is a default placeholder for any unknown type for the renum.ErrTest enum.
	ErrTestUndefinedEnumValue ErrTest = iota

	// ErrTestFoo is an enum value for type ErrTest.
	// ErrTestFoo is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.
	ErrTestFoo

	// ErrTestBar is an enum value for type ErrTest.
	// ErrTestBar is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.
	ErrTestBar

	// ErrTestBaz is an enum value for type ErrTest.
	// ErrTestBaz is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.
	ErrTestBaz

	_ErrTestNamespace = `github.com.gen0cide.renum`
	_ErrTestPkgName   = `renum`
	_ErrTestPkgPath   = `github.com/gen0cide/renum`
)

var (
	// ErrUndefinedErrTestEnumValue is thrown when ParseErrTest(s string) cannot locate a valid enum for the provided string.
	ErrUndefinedErrTestEnumValue = errors.New("cannot identify enum for provided value")
)

const _ErrTestName = "undefined_enum_valuefoobarbaz"

const _ErrTestPascalName = "UndefinedEnumValueFooBarBaz"

const _ErrTestCamelName = "undefinedEnumValuefoobarbaz"

const _ErrTestScreamingName = "UNDEFINED_ENUM_VALUEFOOBARBAZ"

const _ErrTestCommandName = "undefined-enum-valuefoobarbaz"

var _ErrTestNames = []string{
	_ErrTestName[0:20],
	_ErrTestName[20:23],
	_ErrTestName[23:26],
	_ErrTestName[26:29],
}

// ErrTestNames returns a list of possible string values of ErrTest.
func ErrTestNames() []string {
	tmp := make([]string, len(_ErrTestNames))
	copy(tmp, _ErrTestNames)
	return tmp
}

var _ErrTestValueSlice = []ErrTest{
	ErrTestUndefinedEnumValue,
	ErrTestFoo,
	ErrTestBar,
	ErrTestBaz,
}

// ErrTestValues returns a list of possible enum values for the ErrTest type.
func ErrTestValues() []ErrTest {
	tmp := make([]ErrTest, len(_ErrTestValueSlice))
	copy(tmp, _ErrTestValueSlice)
	return tmp
}

var _ErrTestValue = map[string]ErrTest{
	_ErrTestName[0:20]:  0,
	_ErrTestName[20:23]: 1,
	_ErrTestName[23:26]: 2,
	_ErrTestName[26:29]: 3,
}

var _ErrTestPascalValue = map[string]ErrTest{
	_ErrTestPascalName[0:18]:  0,
	_ErrTestPascalName[18:21]: 1,
	_ErrTestPascalName[21:24]: 2,
	_ErrTestPascalName[24:27]: 3,
}

var _ErrTestCamelValue = map[string]ErrTest{
	_ErrTestCamelName[0:18]:  0,
	_ErrTestCamelName[18:21]: 1,
	_ErrTestCamelName[21:24]: 2,
	_ErrTestCamelName[24:27]: 3,
}

var _ErrTestScreamingValue = map[string]ErrTest{
	_ErrTestScreamingName[0:20]:  0,
	_ErrTestScreamingName[20:23]: 1,
	_ErrTestScreamingName[23:26]: 2,
	_ErrTestScreamingName[26:29]: 3,
}

var _ErrTestCommandValue = map[string]ErrTest{
	_ErrTestCommandName[0:20]:  0,
	_ErrTestCommandName[20:23]: 1,
	_ErrTestCommandName[23:26]: 2,
	_ErrTestCommandName[26:29]: 3,
}

var _ErrTestMap = map[ErrTest]string{
	0: _ErrTestName[0:20],
	1: _ErrTestName[20:23],
	2: _ErrTestName[23:26],
	3: _ErrTestName[26:29],
}

var _ErrTestPascalMap = map[ErrTest]string{
	0: _ErrTestPascalName[0:18],
	1: _ErrTestPascalName[18:21],
	2: _ErrTestPascalName[21:24],
	3: _ErrTestPascalName[24:27],
}

var _ErrTestCamelMap = map[ErrTest]string{
	0: _ErrTestCamelName[0:18],
	1: _ErrTestCamelName[18:21],
	2: _ErrTestCamelName[21:24],
	3: _ErrTestCamelName[24:27],
}

var _ErrTestScreamingMap = map[ErrTest]string{
	0: _ErrTestScreamingName[0:20],
	1: _ErrTestScreamingName[20:23],
	2: _ErrTestScreamingName[23:26],
	3: _ErrTestScreamingName[26:29],
}

var _ErrTestCommandMap = map[ErrTest]string{
	0: _ErrTestCommandName[0:20],
	1: _ErrTestCommandName[20:23],
	2: _ErrTestCommandName[23:26],
	3: _ErrTestCommandName[26:29],
}

// String implements the Stringer interface.
func (x ErrTest) String() string {
	if str, ok := _ErrTestMap[x]; ok {
		return str
	}

	return _ErrTestMap[ErrTest(0)]
}

// SnakeCase returns the enum as a snake_case string.
func (x ErrTest) SnakeCase() string {
	if str, ok := _ErrTestMap[x]; ok {
		return str
	}

	return _ErrTestMap[ErrTest(0)]
}

// PascalCase returns the enum as a PascalCase string.
func (x ErrTest) PascalCase() string {
	if str, ok := _ErrTestPascalMap[x]; ok {
		return str
	}

	return _ErrTestPascalMap[ErrTest(0)]
}

// CamelCase returns the enum as a cascalCase string.
func (x ErrTest) CamelCase() string {
	if str, ok := _ErrTestCamelMap[x]; ok {
		return str
	}

	return _ErrTestCamelMap[ErrTest(0)]
}

// ScreamingCase returns the enum as a SCREAMING_CASE string.
func (x ErrTest) ScreamingCase() string {
	if str, ok := _ErrTestScreamingMap[x]; ok {
		return str
	}

	return _ErrTestScreamingMap[ErrTest(0)]
}

// CommandCase returns the enum as a command-case string.
func (x ErrTest) CommandCase() string {
	if str, ok := _ErrTestCommandMap[x]; ok {
		return str
	}

	return _ErrTestCommandMap[ErrTest(0)]
}

var _ErrTestKinds = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `renum.ErrTestUndefinedEnumValue`,
	ErrTestFoo:                `renum.ErrTestFoo`,
	ErrTestBar:                `renum.ErrTestBar`,
	ErrTestBaz:                `renum.ErrTestBaz`,
}

// Kind returns a string of the Go type for the given message.
func (x ErrTest) Kind() string {
	if str, ok := _ErrTestKinds[x]; ok {
		return str
	}

	return _ErrTestKinds[ErrTest(0)]
}

var _ErrTestSources = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `github.com/gen0cide/renum.ErrTestUndefinedEnumValue`,
	ErrTestFoo:                `github.com/gen0cide/renum.ErrTestFoo`,
	ErrTestBar:                `github.com/gen0cide/renum.ErrTestBar`,
	ErrTestBaz:                `github.com/gen0cide/renum.ErrTestBaz`,
}

// Source returns an import path directly to the type.
func (x ErrTest) Source() string {
	if str, ok := _ErrTestSources[x]; ok {
		return str
	}

	return _ErrTestSources[ErrTest(0)]
}

var _ErrTestPaths = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `github.com.gen0cide.renum.err_test_undefined_enum_value`,
	ErrTestFoo:                `github.com.gen0cide.renum.err_test_foo`,
	ErrTestBar:                `github.com.gen0cide.renum.err_test_bar`,
	ErrTestBaz:                `github.com.gen0cide.renum.err_test_baz`,
}

// Source returns an import path directly to the type.
func (x ErrTest) Path() string {
	if str, ok := _ErrTestPaths[x]; ok {
		return str
	}

	return _ErrTestPaths[ErrTest(0)]
}

// PackageName returns the name of the parent package for this type.
func (x ErrTest) PackageName() string {
	return _ErrTestPkgName
}

// ImportPath returns the full import path of the parent package
func (x ErrTest) ImportPath() string {
	return _ErrTestPkgPath
}

// Namespace implements the emitter.Namespaced interface.
func (x ErrTest) Namespace() string {
	return _ErrTestNamespace
}

var _ErrTestDescriptions = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `This value is a default placeholder for any unknown type for the renum.ErrTest enum.`,
	ErrTestFoo:                `ErrTestFoo is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.`,
	ErrTestBar:                `ErrTestBar is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.`,
	ErrTestBaz:                `ErrTestBaz is thrown as part of a unit test for the github.com/gen0cide/renum package and should not be used for any other purpose.`,
}

// Description implements the emitter.Detailed interface.
func (x ErrTest) Description() string {
	if str, ok := _ErrTestDescriptions[x]; ok {
		return str
	}

	return _ErrTestDescriptions[ErrTest(0)]
}

var _ErrTestYARPCCodes = map[ErrTest]yarpcerrors.Code{
	ErrTestUndefinedEnumValue: yarpcerrors.CodeInternal,
	ErrTestFoo:                yarpcerrors.CodeInternal,
	ErrTestBar:                yarpcerrors.CodeInternal,
	ErrTestBaz:                yarpcerrors.CodeInternal,
}

// ToYARPC implements the emitter.YARPCResponder interface.
func (x ErrTest) ToYARPC() yarpcerrors.Code {
	if c, ok := _ErrTestYARPCCodes[x]; ok {
		return c
	}

	return yarpcerrors.CodeUnknown
}

// YARPCError implements the yarpcerrors.IsStatus() interface.
func (x ErrTest) YARPCError() *yarpcerrors.Status {
	return yarpcerrors.Newf(x.ToYARPC(), x.Error())
}

var _ErrTestHTTPCodes = map[ErrTest]int{
	ErrTestUndefinedEnumValue: 500,
	ErrTestFoo:                500,
	ErrTestBar:                500,
	ErrTestBaz:                500,
}

// ToHTTP implements the emitter.HTTPResponder interface.
func (x ErrTest) ToHTTP() int {
	if c, ok := _ErrTestHTTPCodes[x]; ok {
		return c
	}

	return 520
}

var _ErrTestOSExitCodes = map[ErrTest]int{
	ErrTestUndefinedEnumValue: 1,
	ErrTestFoo:                1,
	ErrTestBar:                1,
	ErrTestBaz:                1,
}

// ToHTTP implements the emitter.HTTPResponder interface.
func (x ErrTest) ToOSExit() int {
	if c, ok := _ErrTestOSExitCodes[x]; ok {
		return c
	}

	return 1
}

var _ErrTestMessages = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `undefined enum value for type renum.ErrTest`,
	ErrTestFoo:                `renum error of type foo`,
	ErrTestBar:                `renum error of type bar`,
	ErrTestBaz:                `renum error of type baz`,
}

// Message returns the enum short message description.
func (x ErrTest) Message() string {
	if str, ok := _ErrTestMessages[x]; ok {
		return str
	}

	return _ErrTestMessages[ErrTest(0)]
}

var _ErrTestErrors = map[ErrTest]string{
	ErrTestUndefinedEnumValue: `github.com.gen0cide.renum.err_test_undefined_enum_value (0): undefined enum value for type renum.ErrTest`,
	ErrTestFoo:                `github.com.gen0cide.renum.err_test_foo (1): renum error of type foo`,
	ErrTestBar:                `github.com.gen0cide.renum.err_test_bar (2): renum error of type bar`,
	ErrTestBaz:                `github.com.gen0cide.renum.err_test_baz (3): renum error of type baz`,
}

// Error implements the error interface.
func (x ErrTest) Error() string {
	if str, ok := _ErrTestErrors[x]; ok {
		return str
	}

	return _ErrTestErrors[ErrTest(0)]
}

// ParseErrTest attempts to convert a string to a ErrTest
func ParseErrTest(name string) (ErrTest, error) {
	if len(name) < 1 {
		return ErrTest(0), ErrUndefinedErrTestEnumValue
	}

	first, _ := utf8.DecodeRuneInString(name)
	if first == utf8.RuneError {
		return ErrTest(0), ErrUndefinedErrTestEnumValue
	}

	switch {
	case unicode.IsLower(first):
		// test for snake_case
		if x, ok := _ErrTestValue[name]; ok {
			return x, nil
		}

		// test for command-case
		if x, ok := _ErrTestCommandValue[name]; ok {
			return x, nil
		}

		// test for camelCase
		if x, ok := _ErrTestCamelValue[name]; ok {
			return x, nil
		}
	case unicode.IsUpper(first):
		// test for PascalCase
		if x, ok := _ErrTestPascalValue[name]; ok {
			return x, nil
		}

		// test for SCREAMING_CASE
		if x, ok := _ErrTestScreamingValue[name]; ok {
			return x, nil
		}
	default:
		return ErrTest(0), ErrUndefinedErrTestEnumValue
	}

	return ErrTest(0), ErrUndefinedErrTestEnumValue
}

// MarshalText implements the text marshaller method
func (x ErrTest) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ErrTest) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseErrTest(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Scan implements the Scanner interface.
func (x *ErrTest) Scan(value interface{}) error {
	var name string

	switch v := value.(type) {
	case string:
		name = v
	case []byte:
		name = string(v)
	case nil:
		*x = ErrTest(0)
		return nil
	}

	tmp, err := ParseErrTest(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Value implements the driver Valuer interface.
func (x ErrTest) Value() (driver.Value, error) {
	return x.String(), nil
}

// LookupErrTest attempts to convert a int to it's equivelent ErrTest value.
func LookupErrTest(id int) (ErrTest, error) {
	if _, ok := _ErrTestMap[ErrTest(id)]; ok {
		return ErrTest(id), nil
	}
	return ErrTest(0), fmt.Errorf("%T(%v) is not a valid ErrTest, try [%s]", id, id, strings.Join(_ErrTestNames, ", "))
}

// Code implements the Coder interface.
func (x ErrTest) Code() int {
	return int(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (x ErrTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *ErrTest) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("error unmarshaling JSON value: %v", err)
	}
	tmp, err := ParseErrTest(s)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (x ErrTest) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (x *ErrTest) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return fmt.Errorf("error unmarshaling YAML value: %v", err)
	}

	tmp, err := ParseErrTest(s)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
