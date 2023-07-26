//go:generate renum -c test_errors.yml generate -o . --go-filename "generated_errs_test.go"

package renum

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"

	"golang.org/x/xerrors"
)

var terr = xerrors.New("basic testing error")

var twrappedBasic = Wrap(ErrTestFoo, terr)

var tmiddle = xerrors.Errorf("wrapping middleware: %w", twrappedBasic)

var tmiddle2 = errors.Wrap(tmiddle, "pkg/errors wrapped")

var doubleWrapped = Wrap(ErrTestBaz, tmiddle2)

func TestWrap(t *testing.T) {
	type args struct {
		e   Error
		err error
	}
	tests := []struct {
		name string
		args args
		want Wrapped
	}{
		{
			name: "foo",
			args: args{
				e:   ErrTestFoo,
				err: terr,
			},
			want: &wrapped{
				EmbeddedError: ErrTestFoo,
				Attachment:    terr,
			},
		},
		{
			name: "bar",
			args: args{
				e:   ErrTestBar,
				err: terr,
			},
			want: &wrapped{
				EmbeddedError: ErrTestBar,
				Attachment:    terr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.e, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapped_Typed(t *testing.T) {
	type fields struct {
		Err        Error
		Attachment error
	}
	tests := []struct {
		name   string
		fields fields
		want   Error
	}{
		{
			name: "foo",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			want: ErrTestFoo,
		},
		{
			name: "bar",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: terr,
			},
			want: ErrTestBar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wrapped{
				EmbeddedError: tt.fields.Err,
				Attachment:    tt.fields.Attachment,
			}
			if got := w.Typed(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrapped.Typed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapped_Unwrap(t *testing.T) {
	type fields struct {
		Err        Error
		Attachment error
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "simple",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			want: terr,
		},
		{
			name: "nested",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			want: twrappedBasic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wrapped{
				EmbeddedError: tt.fields.Err,
				Attachment:    tt.fields.Attachment,
			}
			if err := w.Unwrap(); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("wrapped.Unwrap() error = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_wrapped_Is(t *testing.T) {
	type fields struct {
		Err        Error
		Attachment error
	}
	type args struct {
		e error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		deep   bool
	}{
		{
			name: "simple-enum-comparison",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			args: args{
				e: ErrTestFoo,
			},
			want: true,
			deep: true,
		},
		{
			name: "simple-enum-comparison-negative-test",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			args: args{
				e: ErrTestBar,
			},
			want: false,
			deep: false,
		},
		{
			name: "simple-wrapped",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			args: args{
				e: twrappedBasic,
			},
			want: true,
			deep: true,
		},
		{
			name: "simple-wrapped-negative-test",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: terr,
			},
			args: args{
				e: twrappedBasic,
			},
			want: false,
			deep: false,
		},
		{
			name: "nested-wrapped",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			args: args{
				e: ErrTestFoo,
			},
			want: false,
			deep: true,
		},
		{
			name: "nested-wrapped-negative-test",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: twrappedBasic,
			},
			args: args{
				e: ErrTestBar,
			},
			want: false,
			deep: false,
		},
		{
			name: "nested-wrapped-original",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			args: args{
				e: terr,
			},
			want: false,
			deep: true,
		},
		{
			name: "nested-wrapped-original-negative-test",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			args: args{
				e: xerrors.New("not real"),
			},
			want: false,
			deep: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wrapped{
				EmbeddedError: tt.fields.Err,
				Attachment:    tt.fields.Attachment,
			}
			if got := w.Is(tt.args.e); got != tt.want {
				t.Errorf("wrapped.Is() = %v, want %v", got, tt.want)
			}
			if deepgot := xerrors.Is(w, tt.args.e); deepgot != tt.deep {
				t.Errorf("(deep test) xerrors.Is() = %v, want %v", deepgot, tt.deep)
			}
		})
	}
}

func Test_wrapped_Cause(t *testing.T) {
	type fields struct {
		Err        Error
		Attachment error
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "simple",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: terr,
			},
			want: terr,
		},
		{
			name: "nested",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			want: twrappedBasic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wrapped{
				EmbeddedError: tt.fields.Err,
				Attachment:    tt.fields.Attachment,
			}
			if err := w.Cause(); err != tt.want {
				t.Errorf("wrapped.Cause() error = %v, want %v", err, tt.want)
			}
		})
	}
}

func Test_wrapped_Errors(t *testing.T) {
	type fields struct {
		Err        Error
		Attachment error
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
		match  bool
	}{
		{
			name: "simple",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: ErrTestBar,
			},
			want: []error{
				ErrTestFoo,
				ErrTestBar,
			},
			match: true,
		},
		{
			name: "simple-negative-out-of-order",
			fields: fields{
				Err:        ErrTestFoo,
				Attachment: ErrTestBar,
			},
			want: []error{
				ErrTestBar,
				ErrTestFoo,
			},
			match: false,
		},
		{
			name: "wrapped",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: twrappedBasic,
			},
			want: []error{
				ErrTestBar,
				ErrTestFoo,
				terr,
			},
			match: true,
		},
		{
			name: "double-wrapped",
			fields: fields{
				Err:        ErrTestBar,
				Attachment: doubleWrapped,
			},
			want: []error{
				ErrTestBar,
				ErrTestBaz,
				ErrTestFoo,
				terr,
			},
			match: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wrapped{
				EmbeddedError: tt.fields.Err,
				Attachment:    tt.fields.Attachment,
			}
			if got := w.Errors(); reflect.DeepEqual(got, tt.want) != tt.match {
				t.Errorf("wrapped.Errors() = %v, want %v (should match = %v)", got, tt.want, tt.match)
			}
		})
	}
}
