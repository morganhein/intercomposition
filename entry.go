package intercomposition

import (
	"github.com/morganhein/intercomposition/formats"
	"github.com/morganhein/intercomposition/formats/rpm"
)

func NewBuilder() formats.Builder {
	return &rpm.RPM{BaseBuilder: &formats.BaseBuilder{}}
}
