package renum

// Typer can be implemented by types to allow their callers to
// understand their location within Go source.
type Typer interface {
	Kind() string
	Type() string
	Source() string
	PackageName() string
	ImportPath() string
}
