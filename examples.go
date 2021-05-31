package daterange

import (
	"fmt"
	"time"
)

func main() {
	event, _ := New("2021-01-01T05:00:00Z", "2021-01-02T08:00:00Z")
	event2, err := New("2021-01-02T06:00:00Z", "2021-01-03T08:00:00Z")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(event)

	// get duration of time range (end - start)
	fmt.Println(event.Duration())

	// check if two time ranges intersect
	fmt.Println("Intersect?", event.Intersect(*event2))

	// check if two time ranges are equivalent
	event.Equal(*event2)

	// lengthen ending time by 3 hours
	event.Lengthen(time.Hour * 3)

	// shorten ending time by 3 hours
	event.Lengthen(time.Hour * 3)

	// shorten ending time by 48 hours
	// if shorten duration is longer than entire event, ending time is set to start
	event.Lengthen(time.Hour * 48)
}
