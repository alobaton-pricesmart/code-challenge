package flights

import (
	"errors"
	"sort"
)

// Flights Struct to map the employee flights.
type Flights struct {
	Costs [][]int
}

// Process Process the min costs throught employee flights.
func (f *Flights) MinCost() (int, error) {
	err := f.validate()
	if err != nil {
		return -1, err
	}

	min := 0
	var differences []int

	for _, cost := range f.Costs {
		min += cost[0]
		diff := cost[1] - cost[0]
		differences = append(differences, diff)
	}

	sort.Ints(differences)

	for i := 0; i < len(f.Costs)/2; i++ {
		min += differences[i]
	}

	return min, nil
}

// validate Private function that validates the struct
func (f *Flights) validate() error {
	if f.Costs == nil || len(f.Costs) == 0 {
		return errors.New("Costs can't be nil or empty")
	}

	if len(f.Costs)%2 != 0 {
		return errors.New("Costs len must be even")
	}

	for _, cost := range f.Costs {
		if len(cost) != 2 {
			return errors.New("cost len must be 2")
		}
	}

	return nil
}
