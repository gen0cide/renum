package renum

// Sourcer requires enums to be able to self describe aspects of the Go source and package
// which they're located. This makes Enum's great for tracing and error handling. This
// interface allows callers to retrieve additional context of the Enum value without having
// to take up excess memory space, since the enum is just a type alias to a builtin numeric.
type Sourcer interface {
	PackageName() string // foo
	PackagePath() string // github.com/gen0cide/foo
	ExportType() string  // foo.TypeVal
	ExportRef() string   // github.com/gen0cide/foo.TypeVal
}
