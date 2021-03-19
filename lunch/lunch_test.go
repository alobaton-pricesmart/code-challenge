package lunch

import (
	"testing"
)

func TestLunchCase1(t *testing.T) {
	n := 4
	boxes := []int{0, 0, 1, 0 }
	preferences := []int{1, 0, 0, 0}
	l := Lunch{N: n, Boxes: boxes, Preferences: preferences}

	got, err := l.QueueLen()
	if err != nil {
		t.Errorf("Error executing f.Process")
	}

	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLunchCase2(t *testing.T) {
	n := 4
	boxes := []int{0, 1, 0, 0 }
	preferences := []int{0, 0, 0, 0}
	l := Lunch{N: n, Boxes: boxes, Preferences: preferences}

	got, err := l.QueueLen()
	if err != nil {
		t.Errorf("Error executing f.Process")
	}

	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}