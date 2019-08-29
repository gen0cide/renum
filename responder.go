package renum

import "go.uber.org/yarpc/yarpcerrors"

// YARPCResponder allows a type to define a specified YARPC error code so that a handler can
// automatically act on it's behalf without having to maintain a separate mapping.
type YARPCResponder interface {
	ToYARPC() yarpcerrors.Code
	YARPCError() *yarpcerrors.Status
}

// HTTPResponder allows a type to define a specified HTTP status code so that a handler
// can automatically act on it's behalf without having to maintain a separate mapping.
type HTTPResponder interface {
	ToHTTP() int
}

// ProcessResponder allows a type to define the exit code that a process should exit with
// upon encountering said type. This primarily targets error handling.
type ProcessResponder interface {
	ToOSExit() int
}
