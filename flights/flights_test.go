package flights

import (
	"testing"
)

func TestGetMinCostDefault(t *testing.T) {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	f := Flights{Costs: costs}

	err := f.Process()
	if err != nil {
		t.Errorf("Error executing f.Process")
	}

	got := f.MinCost()
	want := 110

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetMinCost(t *testing.T) {
	costs := [][]int{{10, 20}, {30, 200}, {50, 400}, {30, 20}}
	f := Flights{Costs: costs}

	err := f.Process()
	if err != nil {
		t.Errorf("Error executing f.Process")
	}

	got := f.MinCost()
	want := 120

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestGetMinCostFailNil(t *testing.T) {
	f := Flights{}

	err := f.Process()
	if err == nil {
		t.Errorf("err can't be nil")
	}
}

func TestGetMinCostFailOdd(t *testing.T) {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}}
	f := Flights{Costs: costs}

	err := f.Process()
	if err == nil {
		t.Errorf("err can't be nil")
	}
}

func TestGetMinCostFailCostOdd(t *testing.T) {
	costs := [][]int{{10, 20}, {30}}
	f := Flights{Costs: costs}

	err := f.Process()
	if err == nil {
		t.Errorf("err can't be nil")
	}
}
