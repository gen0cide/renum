package renum

import (
	"encoding/json"
	"fmt"

	"go.uber.org/yarpc/yarpcerrors"

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

	// Cause implements github.com/pkg/errors.Causer interface
	Cause() error

	// Unwrap implements the golang.org/x/xerrors.Wrapper interface.
	Unwrap() error

	// Is implements the golang.org/x/xerrors.Is interface.
	Is(e error) bool

	// Format implements fmt.Formatter interface (old error handling)
	Format(f fmt.State, c rune)

	// FormatError implements golang.org/x/xerrors.Formatter interface (new error handling)
	FormatError(p xerrors.Printer) error

	// Errors implements the github.com/uber-go/multierr.errorGroup interface.
	Errors() []error

	// YARPCError implements the go.uber.org/yarpc/yarpcerrors interface for creating
	// custom YARPC errors.
	YARPCError() *yarpcerrors.Status
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

// Cause implements github.com/pkg/errors.Causer interface.
func (w *wrapped) Cause() error {
	return w.Attachment
}

// Format implements fmt.Formatter
func (w *wrapped) Format(f fmt.State, c rune) {
	xerrors.FormatError(w, f, c)
}

// FormatError implements xerrors.Formatter
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

var _builtinYARPCCode = yarpcerrors.CodeUnknown

// YARPCError is used by YARPC to ensure a wrapped error retains wrapped context information
// as the error is returned to the caller across the network. YARPCError attempt to extract the
// first yarpcerrors.Code (provided one of the errors if of type renum.YARPCError) and override
// the default (2=CodeUnknown). It will used the renum.Wrapped Typed() (top level renum.Error)
// for the message and name fields, with the name field being displayed in command-case (as per spec).
// The resulting structure will look like this on the wire:
//
//  {
//  	"code" 2,
//  	"name": "cant-do-that",
//  	"message": "example.namespace.path.err_cant_do_that (2): example error message",
//  	"details": "[...JSON string...]"
//  }
//
// You can use yarpcerrors.FromError(err error) to turn your returned RPC caller's error back into
// a *yarpcerrors.Status, and subsequently pass that to ExtractErrorsFromYARPCStatus to get
// the detailed information that was embedded in the details field.
func (w *wrapped) YARPCError() *yarpcerrors.Status {
	errs := w.Errors()
	errcode := _builtinYARPCCode
	typedErr := w.Typed()
	for _, x := range errs {
		if yerr, ok := x.(YARPCError); ok {
			errcode = yerr.ToYARPC()
			break
		}
	}

	ret := yarpcerrors.Newf(errcode, "%s", typedErr.Error())
	if ret == nil {
		return nil
	}

	detailBytes, err := json.Marshal(extractTypeInfoFromList(errs...))
	if err != nil {
		return ret
	}

	return ret.WithDetails(detailBytes)
}

// ExtractErrorsFromYARPCStatus is a helper method to read in a YARPC Status has been transmitted
// across application boundries and attempts to unpack an error stack of the foreign services
// wrapped errors. Note that this should *not* be used to programmatically type check errors, but
// rather in presenting the remote error's contexts in ways that are easily formatted for a user.
func ExtractErrorsFromYARPCStatus(status *yarpcerrors.Status) ([]ErrorTypeInfo, bool) {
	var ret []ErrorTypeInfo
	if status == nil {
		return ret, false
	}

	if status.Details() == nil {
		return ret, false
	}

	err := json.Unmarshal(status.Details(), &ret)
	if err != nil {
		return ret, false
	}

	return ret, true
}
