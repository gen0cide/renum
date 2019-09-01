package renum

// EnumTypeInfo is a type used to hold all the metadata associated with a given renum.Enum where
// the fields of this structure are associated directly with the return values from the renum.Enum interface.
// This acts as a convenience to help things like structured loggers or HTTP JSON responses to be have
// information extracted into a self contained object.
type EnumTypeInfo struct {
	Name    string      `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Code    int         `json:"code,omitempty" mapstructure:"code,omitempty" yaml:"code,omitempty" toml:"code,omitempty"`
	Details TypeDetails `json:"details,omitempty" mapstructure:"details,omitempty" yaml:"details,omitempty" toml:"details,omitempty"`
}

// ExtractEnumTypeInfo is used to take a renum.Enum type and expand it's details into a more annotated
// structure. The primary purpose of this is to act as a helper to loggers who wish to expand interface methods
// of the renum.Enum type into a nested, flat structure.
func ExtractEnumTypeInfo(e Enum) EnumTypeInfo {
	return EnumTypeInfo{
		// Name: e.String(),
		// Code: e.Code(),
		// Details: TypeDetails{
		// 	Namespace:   e.Namespace(),
		// 	Path:        e.Path(),
		// 	Kind:        e.Kind(),
		// 	Source:      e.Source(),
		// 	ImportPath:  e.ImportPath(),
		// 	Description: e.Description(),
		// },
	}
}

// TypeDetails allows for detailed renum.Enum/renum.Error type information to be embedded in a nested structure
// for TypeInfo structs.
type TypeDetails struct {
	Namespace   string `json:"namespace,omitempty" mapstructure:"namespace,omitempty" yaml:"namespace,omitempty" toml:"namespace,omitempty"`
	Path        string `json:"path,omitempty" mapstructure:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	Kind        string `json:"kind,omitempty" mapstructure:"kind,omitempty" yaml:"kind,omitempty" toml:"kind,omitempty"`
	Source      string `json:"source,omitempty" mapstructure:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty"`
	ImportPath  string `json:"import_path,omitempty" mapstructure:"import_path,omitempty" yaml:"import_path,omitempty" toml:"import_path,omitempty"`
	Description string `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
}

// ErrorTypeInfo is a type used to hold all the metadata associated with a given renum.Error where
// the fields of this structure are associated directly with the return values from the renum.Error interface.
// This acts as a convenience to help things like structured loggers or HTTP JSON responses to be have
// information extracted into a self contained object.
type ErrorTypeInfo struct {
	Name    string      `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Code    int         `json:"code,omitempty" mapstructure:"code,omitempty" yaml:"code,omitempty" toml:"code,omitempty"`
	Details TypeDetails `json:"details,omitempty" mapstructure:"details,omitempty" yaml:"details,omitempty" toml:"details,omitempty"`
	Message string      `json:"message,omitempty" mapstructure:"message,omitempty" yaml:"message,omitempty" toml:"message,omitempty"`
}

// ExtractErrorTypeInfo is used to take a renum.Error type and expand it's details into a more annotated
// structure. The primary purpose of this is to act as a helper to loggers who wish to expand interface methods
// of the renum.Error type into a nested, flat structure.
func ExtractErrorTypeInfo(e error) []ErrorTypeInfo {
	if werr, ok := e.(Wrapped); ok {
		return extractTypeInfoFromList(werr.Errors()...)
	}

	ret := []ErrorTypeInfo{}

	if rerr, ok := e.(Error); ok {
		ret = append(ret, typedErrorToInfo(rerr))
		return ret
	}

	return append(ret, basicErrorToInfo(e))
}

// extractTypeInfoFromList does the heavy work of assigning a list of errors
// into their appropriate casts.
func extractTypeInfoFromList(errs ...error) []ErrorTypeInfo {
	ret := make([]ErrorTypeInfo, len(errs))

	for idx, err := range errs {
		if rerr, ok := err.(Error); ok {
			ret[idx] = typedErrorToInfo(rerr)
			continue
		}
		ret[idx] = basicErrorToInfo(err)
	}

	return ret
}

// typedErrorToInfo converts a renum.Error type into an ErrorTypeInfo structure.
func typedErrorToInfo(e Error) ErrorTypeInfo {
	return ErrorTypeInfo{
		// Name: e.String(),
		// Code: e.Code(),
		// Details: TypeDetails{
		// 	Namespace:   e.Namespace(),
		// 	Path:        e.Path(),
		// 	Kind:        e.Kind(),
		// 	Source:      e.Source(),
		// 	ImportPath:  e.ImportPath(),
		// 	Description: e.Description(),
		// },
		Message: e.Message(),
	}
}

// basicErrorToInfo converts a basic error type into an ErrorTypeInfo structure.
func basicErrorToInfo(e error) ErrorTypeInfo {
	if e == nil {
		return ErrorTypeInfo{}
	}

	return ErrorTypeInfo{
		Message: e.Error(),
	}
}
