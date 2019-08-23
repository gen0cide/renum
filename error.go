package renum

// Error allows types to conform to a strongly defined interface, as well as act as enriched error builtins.
// The point of this is that as types that satisfy Error pass across package boundry, context and metadata
// is not lost.
type Error interface {
	Enum
	error
}

// YARPCError extends the Error interface to allow a type to additionally self-report a YARPC error code
// in order to enrich the handler's ability to respond with the proper code when an error of this type
// is encountered.
type YARPCError interface {
	Error
	YARPCResponder
}

// HTTPError extends the Error interface to allow a type to additionally self-report an HTTP Response code
// in order to enrich a net/http handler's ability to respond with the proper status code when an error
// of this type is encountered.
type HTTPError interface {
	Error
	HTTPResponder
}

// ProcessError extends the Error interface to allow a type to additionally self-report a specific
// exit code it wishes the handler to exit the process with.
type ProcessError interface {
	Error
	ProcessResponder
}
