package renum

// Namespacer requires a type be able to produce information relating to
// the package or component it's defined by. This allows tracing of errors
// to propogate across package boundries without loosing the ability to easily
// identify the owner of a type.
type Namespacer interface {
	Path() string
	Namespace() string
}
