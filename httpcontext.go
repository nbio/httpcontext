package httpcontext

import (
	"errors"
	"io"
	"net/http"
)

// ContextReadCloser implements the io.ReadCloser interface
// with two additional methods: Context() and SetContext().
type ContextReadCloser interface {
	io.ReadCloser
	Context() interface{}
	SetContext(interface{})
}

type contextReadCloser struct {
	io.ReadCloser
	context interface{}
}

func (crc *contextReadCloser) Context() interface{} {
	return crc.context
}

func (crc *contextReadCloser) SetContext(context interface{}) {
	crc.context = context
}

// Set sets a context value on req.
// It accomplishes this by replacing the http.Request.Body with
// a ContextReadCloser. See Invasion of the Body Snatchers.
func Set(req *http.Request, context interface{}) {
	crc, ok := req.Body.(ContextReadCloser)
	if !ok {
		crc = &contextReadCloser{ReadCloser: req.Body}
		req.Body = crc
	}
	crc.SetContext(context)
}

// Set gets a context value from req.
// Set will return an error if req.Body was not previously
// replaced with a ContextReadCloser.
func Get(req *http.Request) (interface{}, error) {
	crc, ok := req.Body.(ContextReadCloser)
	if !ok {
		return nil, errors.New("Unable to convert http.Request.Body to ContextReadCloser")
	}
	return crc.Context(), nil
}
