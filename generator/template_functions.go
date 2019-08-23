package generator

import (
	"fmt"
)

var (
	_builtinYARPCCode  = 2
	_builtinHTTPCode   = 520
	_builtinOSExitCode = 1
	_yarpcCodes        = map[int]string{
		0:  "yarpcerrors.CodeOK",
		1:  "yarpcerrors.CodeCancelled",
		2:  "yarpcerrors.CodeUnknown",
		3:  "yarpcerrors.CodeInvalidArgument",
		4:  "yarpcerrors.CodeDeadlineExceeded",
		5:  "yarpcerrors.CodeNotFound",
		6:  "yarpcerrors.CodeAlreadyExists",
		7:  "yarpcerrors.CodePermissionDenied",
		8:  "yarpcerrors.CodeResourceExhausted",
		9:  "yarpcerrors.CodeFailedPrecondition",
		10: "yarpcerrors.CodeAborted",
		11: "yarpcerrors.CodeOutOfRange",
		12: "yarpcerrors.CodeUnimplemented",
		13: "yarpcerrors.CodeInternal",
		14: "yarpcerrors.CodeUnavailable",
		15: "yarpcerrors.CodeDataLoss",
		16: "yarpcerrors.CodeUnauthenticated",
	}
	_finishBrace = `}`
)

// DefaultYARPCCode is a template helper that returns the default YARPC code (2, CodeUnknown).
func DefaultYARPCCode() int {
	return _builtinYARPCCode
}

// DefaultHTTPCode is a template helper that returns the default HTTP status code (520, Error Unknown).
func DefaultHTTPCode() int {
	return _builtinHTTPCode
}

// DefaultOSExitCode is a template helper function that returns the default OS Exit code (1, errored).
func DefaultOSExitCode() int {
	return _builtinOSExitCode
}

// UnderscoreStringify returns a string that is all of the enum value names concatenated without a separator, in snake_case.
func UnderscoreStringify(c Config) (ret string, err error) {
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret += val.Snake()
		}
	}
	return
}

// PascalizeStringify returns a string that is all of the enum value names as Pascalized strings concatenated without a separator.
func PascalizeStringify(c Config) (ret string, err error) {
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret += val.Pascal()
		}
	}
	return
}

// MapifyDescriptions creates a templated map of all enum values and their associated descriptions.
func MapifyDescriptions(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.EnumID())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s`,\n", ret, val.PrefixedPascal(), val.Description)
		}
	}
	ret += _finishBrace
	return
}

// MapifyMessages creates a templated map of all enum values and their associated message values.
func MapifyMessages(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s`,\n", ret, val.PrefixedPascal(), val.Message)
		}
	}
	ret += _finishBrace
	return
}

// MapifyKinds creates a templated map of all enum values and their kind strings.
func MapifyKinds(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s.%s`,\n", ret, val.PrefixedPascal(), c.Go.PackageName, val.PrefixedPascal())
		}
	}
	ret += _finishBrace
	return
}

// MapifyFullKinds creates a templated map of all enum values and a detailed kind string.
func MapifyFullKinds(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s.%s`,\n", ret, val.PrefixedPascal(), c.Go.PackagePath, val.PrefixedPascal())
		}
	}
	ret += _finishBrace
	return
}

// MapifyFullPaths creates a templated map of all enum values and a complete namespace and type string.
func MapifyFullPaths(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s.%s_%s`,\n", ret, val.PrefixedPascal(), c.Plugins.Namespace.Namespace, c.Go.Prefix.Snake(), val.Snake())
		}
	}
	ret += _finishBrace
	return
}

// MapifyErrorMessages creates a templated map of all enum values and their pre-rendered error messages.
func MapifyErrorMessages(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: `%s.%s_%s (%d): %s`,\n", ret, val.PrefixedPascal(), c.Plugins.Namespace.Namespace, c.Go.Prefix.Snake(), val.Snake(), val.Value, val.Message)
		}
	}
	ret += _finishBrace
	return
}

// MapifyHTTPCodes creates a templated map of all enum values and their associated HTTP status codes.
func MapifyHTTPCodes(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]int{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: %d,\n", ret, val.PrefixedPascal(), val.HTTPCode())
		}
	}
	ret += _finishBrace
	return
}

// MapifyOSExitCodes creates a templated map of all enum values and their associated os exit codes.
func MapifyOSExitCodes(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]int{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: %d,\n", ret, val.PrefixedPascal(), val.OSExitCode())
		}
	}
	ret += _finishBrace
	return
}

// MapifyYARPCCodes creates a templated map of all enum values and their associated yarpcerrors.Code values.
func MapifyYARPCCodes(c Config) (ret string, err error) {
	ret = fmt.Sprintf("map[%s]yarpcerrors.Code{\n", c.Go.Prefix.Pascal())
	for _, val := range c.Values {
		yarpcVal, ok := _yarpcCodes[val.YARPCCode()]
		if !ok {
			err = fmt.Errorf("could not find yarpc code for %v (code=%d)", val.PrefixedPascal(), val.YARPCCode())
			return
		}
		if val.Pascal() != skipHolder {
			ret = fmt.Sprintf("%s%s: %s,\n", ret, val.PrefixedPascal(), yarpcVal)
		}
	}
	ret += _finishBrace
	return
}

// Mapify returns a map that is all of the indexes for a string value lookup
func Mapify(c Config) (ret string, err error) {
	strName := fmt.Sprintf(`_%sName`, c.Go.Prefix.Pascal())
	ret = fmt.Sprintf("map[%s]string{\n", c.Go.Prefix.Pascal())
	index := 0
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			nextIndex := index + len(val.Snake())
			ret = fmt.Sprintf("%s%d: %s[%d:%d],\n", ret, val.Value, strName, index, nextIndex)
			index = nextIndex
		}
	}
	ret += _finishBrace
	return
}

// Unmapify returns a map that is all of the indexes for a string value lookup
func Unmapify(c Config, lowercase bool) (ret string, err error) {
	strName := fmt.Sprintf(`_%sName`, c.Go.Prefix.Pascal())
	ret = fmt.Sprintf("map[string]%s{\n", c.Go.Prefix.Pascal())
	index := 0
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			nextIndex := index + len(val.Snake())
			ret = fmt.Sprintf("%s%s[%d:%d]: %d,\n", ret, strName, index, nextIndex, val.Value)
			if lowercase {
				ret = fmt.Sprintf("%sstrings.ToLower(%s[%d:%d]): %d,\n", ret, strName, index, nextIndex, val.Value)
			}
			index = nextIndex
		}
	}
	ret += _finishBrace
	return
}

// Namify returns a slice that is all of the possible names for an enum in a slice
func Namify(c Config) (ret string, err error) {
	strName := fmt.Sprintf(`_%sName`, c.Go.Prefix.Pascal())
	ret = "[]string{\n"
	index := 0
	for _, val := range c.Values {
		if val.Pascal() != skipHolder {
			nextIndex := index + len(val.Snake())
			ret = fmt.Sprintf("%s%s[%d:%d],\n", ret, strName, index, nextIndex)
			index = nextIndex
		}
	}
	ret += _finishBrace
	return
}
