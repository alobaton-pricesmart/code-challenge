package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
)

// Flights Struct to map the employee flights.
type Flights struct {
	Employees []Employee `json:"employees"`
}

// Employee Struct to map the flight costs of an employee.
type Employee struct {
	ID    string `json:"id"`
	Costs []Cost `json:"costs"`
}

// Cost struct to map the cost of flying to a city.
type Cost struct {
	CityCode string `json:"cityCode"`
	Cost     int    `json:"cost"`
}

// ByCost implements sort.Interface in order to get the min cost of a city flights.
type ByCost []Cost

func (c ByCost) Len() int           { return len(c) }
func (c ByCost) Less(i, j int) bool { return c[i].Cost < c[j].Cost }
func (c ByCost) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// main Main function of the program.
func main() {
	body, err := ioutil.ReadFile("input.json")
	handleError(err)

	// Just for development purpose...
	// fmt.Println(string(body))

	flights := Flights{}
	err = json.Unmarshal([]byte(body), &flights)
	handleError(err)

	var minCost int
	minCost, err = getMinCost(flights)
	handleError(err)

	fmt.Printf("%d\n", minCost)
}

// handleError Handle error
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// getMinCost Returns the min costs throught employee flights.
func getMinCost(flights Flights) (int, error) {
	if flights.Employees == nil || len(flights.Employees) == 0 {
		return -1, errors.New("employees can't be nil or empty")
	}

	min := 0
	for _, employee := range flights.Employees {
		if employee.Costs == nil || len(employee.Costs) == 0 {
			return -1, errors.New("costs can't be nil or empty")
		}

		sort.Sort(ByCost(employee.Costs))

		minCost := employee.Costs[0]

		min += minCost.Cost
	}

	return min, nil
}
