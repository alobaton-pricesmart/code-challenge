package main

import (
	"code-challenge/flights"
	"log"
)

// main Main function of the program.
func main() {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	f := flights.Flights{Costs: costs}

	err := f.Process()
	if err != nil {
		log.Fatal("Error executing GetMinCost:", err)
		return
	}

	minCost := f.MinCost()
	log.Print(minCost)
}
