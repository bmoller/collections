// ©2002 Brandon Moller

package collections_test

import (
	"fmt"
	"testing"

	"github.com/bmoller/collections"
)

func TestErrIndexOutOfRange(t *testing.T) {
	err := collections.ErrIndexOutOfRange{
		Index: 100,
		Size:  10,
	}

	if err.Error() != fmt.Sprintf("index %d is invalid for list of length %d", 100, 10) {
		t.Fatalf("unexpected error string: %s", err)
	}
}

func TestErrInvalidRange(t *testing.T) {
	err := collections.ErrInvalidRange{
		End:   10,
		Start: 20,
	}

	if err.Error() != fmt.Sprintf("invalid range with start %d and end %d; valid ranges follow 0 <= start <= end", 20, 10) {
		t.Fatalf("unexpected error string: %s", err)
	}
}
