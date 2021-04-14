package lunch

type Lunch struct {
	N           int
	Boxes       []int
	Preferences []int
}

func (l Lunch) QueueLen() (int, error) {
	queue := make(chan int, l.N)

	boxIndex := 0
	go queuePreferences(queue, &boxIndex, l.Boxes, l.Preferences)

	preferences := []int{}
	for preference := range queue {
		preferences = append(preferences, preference)
	}

	queue = make(chan int, len(preferences))
	go queuePreferences(queue, &boxIndex, l.Boxes, preferences)

	return len(queue), nil
}

func queuePreferences(queue chan int, boxIndex *int, boxes []int, preferences []int) {
	for _, preference := range preferences {
		box := boxes[*boxIndex]

		if box == preference {
			*boxIndex++
		} else {
			queue <- preference
		}
	}
	close(queue)
}
