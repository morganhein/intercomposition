package formats

import (
	"io"
	"net/http"
)

// all builders (apk, deb, rpm) will implement this
type Builder interface {
	HasThis() bool
	OrThis() error
	Setup() error                                           //setup the environment
	Build(buildRequest string, output io.WriteCloser) error //build the package
	Hash(string) (string, error)                            //common functions, whatever
	Extend(string) Builder
}

// and will share a common set of underlying data
// and a base implementation of often repeated functionality
// will keep repetitive operations DRY
type BaseBuilder struct {
	Getter http.Client
}

// BaseBuilder operations on the builder data can be pre-defined and shared
func (b *BaseBuilder) DoSomet() error {
	return nil
}

// and some functionality to implement the Builder interface
// can be predefined
func (b *BaseBuilder) HasThis() bool {
	return true
}

func (b *BaseBuilder) OrThis() error {
	return nil
}

func (b *BaseBuilder) Setup() error {
	return nil
}

func (b *BaseBuilder) Hash(f string) (string, error) {
	return "", nil
}

// but don't implement all functionality, which forces each
// package format to implement these functions or it it won't compile
//func (b *BaseBuilder) Extend(pkg string) Builder {
//	return b
//}
//
//func (b *BaseBuilder) Build(buildRequest string, output io.WriteCloser) error {
//	return nil
//}
