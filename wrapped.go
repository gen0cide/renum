package renum

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ensure the wrapped unexported type conforms to the Wrapped interface.
var _ Wrapped = (*wrapped)(nil)

// EmbeddedError is a type alias for renum.Error that allows the renum.Wrapped interface
// to directly embed the EmbeddedError into Wrapped without conflicting with the
// Go built-in error interface's Error() string method.
type EmbeddedError Error

// Wrap combines a renum.Error type as well as a standard library error in order to allow for
// contextual information.
func Wrap(e Error, err error) Wrapped {
	return &wrapped{
		EmbeddedError: e,
		Attachment:    err,
	}
}

// Wrapped defines a type implementation that allows for wrapped Errors to be wrapped
// and unwrapped following Go convention (both old and new).
type Wrapped interface {
	EmbeddedError

	// Typed returns the renum.Err typed error for this wrapped error.
	Typed() Error

	// implements github.com/pkg/errors.Causer interface
	Cause() error

	// implements the golang.org/x/xerrors.Wrapper interface.
	Unwrap() error

	// implements the golang.org/x/xerrors.Is interface.
	Is(e error) bool

	// implements fmt.Formatter interface (old error handling)
	Format(f fmt.State, c rune)

	// implements golang.org/x/xerrors.Formatter interface (new error handling)
	FormatError(p xerrors.Printer) error

	// implements the github.com/uber-go/multierr.errorGroup interface.
	Errors() []error
}

// WrappedError is used to sidecar a standard library error to a renum.Error in order to
// enrich a renum.Error with additional context.
type wrapped struct {
	EmbeddedError
	Attachment error
}

// Typed implements the renum.Wrapped interface.
func (w *wrapped) Typed() Error {
	return w.EmbeddedError
}

// Unwrap implements the xerrors.Wrapper interface.
func (w *wrapped) Unwrap() error {
	return w.Attachment
}

// Is implements the xerrors interface.
func (w *wrapped) Is(e error) bool {
	if e == nil {
		return false
	}

	if werr, ok := e.(Wrapped); ok {
		return w.Typed() == werr.Typed()
	}

	if rerr, ok := e.(Error); ok {
		return w.Typed() == rerr
	}

	return false
}

// implements github.com/pkg/errors.Causer interface.
func (w *wrapped) Cause() error {
	return w.Attachment
}

// implements fmt.Formatter
func (w *wrapped) Format(f fmt.State, c rune) {
	xerrors.FormatError(w, f, c)
}

// implements xerrors.Formatter
func (w *wrapped) FormatError(p xerrors.Printer) error {
	p.Print(w.Error())
	return nil
}

// Errors handles unwrapping of errors into a stack of []error types
// that is compatible with renum.Wrapped, github.com/pkg/errors, and golang.org/x/xerrors.
func (w *wrapped) Errors() []error {
	ret := []error{w.Typed()}

	// initial error
	next := w.Attachment

	// recursively loop through
	for {
		if next == nil {
			break
		}

		// wrapped error attachment was another wrapped error.
		// unwind the typed error from it and unwrap the next error
		if werr, ok := next.(Wrapped); ok {
			ret = append(ret, werr.Typed())
			next = werr.Unwrap()
			continue
		}

		// wrapped error attachment was an opaque renum.Err.
		// add it to the chain and return
		if rerr, ok := next.(Error); ok {
			ret = append(ret, rerr)
			break
		}

		// handle errors wrapped with xerrors interface (priority)
		if xerr, ok := next.(xerrors.Wrapper); ok {
			next = xerr.Unwrap()
			continue
		}

		// handle errors wrapped with github.com/pkg/errors
		if pkgerr, ok := next.(interface{ Cause() error }); ok {
			next = pkgerr.Cause()
			continue
		}

		ret = append(ret, next)
		break
	}

	return ret
}
