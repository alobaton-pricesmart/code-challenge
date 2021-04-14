package main

import (
	"code-challenge/bintree"
	"code-challenge/flights"
	"code-challenge/lunch"
	"log"

	"golang.org/x/tour/tree"
)

// main Main function of the program.
func main() {

	// flights
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	f := flights.Flights{Costs: costs}

	result, err := f.MinCost()
	if err != nil {
		log.Fatal("Error executing f.MinCost:", err)
		return
	}
	log.Print(result)

	// launch
	n := 4
	boxes := []int{0, 1, 0, 0}
	preferences := []int{0, 0, 0, 0}
	l := lunch.Lunch{N: n, Boxes: boxes, Preferences: preferences}
	result, err = l.QueueLen()
	if err != nil {
		log.Fatal("Error executing l.QueueLen:", err)
		return
	}
	log.Print(result)

	// bintree
	log.Print(bintree.Same(tree.New(1), tree.New(1)))
	log.Print(bintree.Same(tree.New(1), tree.New(2)))
	log.Print(bintree.Same(tree.New(2), tree.New(1)))
}
