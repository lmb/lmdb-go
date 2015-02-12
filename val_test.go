package lmdb

import (
	"testing"
)

func TestVal(t *testing.T) {
	orig := "hey hey"
	val := wrapVal([]byte(orig))

	s := val.String()
	if s != orig {
		t.Errorf("String() not the same as original data: %q", s)
	}

	p := val.Bytes()
	if string(p) != orig {
		t.Errorf("Bytes() not the same as original data: %q", p)
	}
}

func TestValCopy(t *testing.T) {
	orig := "hey hey"
	val := wrapVal([]byte(orig))

	s := val.String()
	if s != orig {
		t.Errorf("String() not the same as original data: %q", s)
	}

	p := val.BytesCopy()
	if string(p) != orig {
		t.Errorf("Bytes() not the same as original data: %q", p)
	}
}
