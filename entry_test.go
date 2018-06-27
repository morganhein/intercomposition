package intercomposition

import (
	"fmt"
	"testing"
)

func TestEntryForRPM(t *testing.T) {
	ba := NewBuilder()
	fmt.Println(ba.HasThis())
	ba = ba.Extend("gcc")
	fmt.Println(ba.HasThis())
}
