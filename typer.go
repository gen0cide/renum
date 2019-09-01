package renum

// Typer can be implemented by types to allow their callers to get a string
// reference to the type that they are. Here's an example of what that means:
//
// 	type Foo int           // enum type alias
// 	const FooValA Foo = 1  // enum value assignment
// 	Foo(1).Type() = "Foo"  // Type() returns the name of the type.
//
type Typer interface {
	Type() string // Type
}
