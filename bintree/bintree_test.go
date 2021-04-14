package bintree

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSameOK(t *testing.T) {
	got := Same(tree.New(1), tree.New(1))

	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSameLeftKO(t *testing.T) {
	got := Same(tree.New(1), tree.New(2))

	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSameRightKO(t *testing.T) {
	got := Same(tree.New(2), tree.New(1))

	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
