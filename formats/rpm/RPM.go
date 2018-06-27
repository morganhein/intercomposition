package rpm

import (
	"github.com/morganhein/intercomposition/formats"
	"io"
)

// Implement each builder type, as normal
// overriding the base Builder functions as necessary
// and optionally including custom data
type RPM struct {
	*formats.BaseBuilder
	rpmSpecificThing string
}

// over-write a single function that will differ
// from the base
func (r *RPM) HasThis() bool {
	return false
}

// implement building
func (r *RPM) Build(buildRequest string, output io.WriteCloser) error {
	return nil
}

// allow hacks to overload any function where necessary to
// support custom functionality for that one package
func (r *RPM) Extend(pkg string) formats.Builder {
	switch pkg {
	case "gcc":
		return &GCCFix{RPM: r}
	}
	return r
}
