package utils

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...any)
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}

func (t *tracer) Trace(a ...any) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct {
}

func (n *nilTracer) Trace(a ...any) {
}

func Off() Tracer {
	return &nilTracer{}
}
