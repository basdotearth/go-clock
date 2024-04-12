package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	usePins := err == nil
	if err != nil {
		fmt.Println(err)
	}

	minutePins := map[int][]int{
		5:  {8, 22},
		10: {23, 22},
		15: {12, 25, 22},
		20: {7, 22},
		25: {7, 8, 22},
		30: {18, 22},
		35: {7, 8, 27},
		40: {7, 27},
		45: {12, 25, 27},
		50: {23, 27},
		55: {8, 27},
		0:  {14},
	}

	hourPins := map[int]int{
		1:  17,
		2:  6,
		3:  0,
		4:  13,
		5:  10,
		6:  5,
		7:  9,
		8:  4,
		9:  11,
		10: 15,
		11: 3,
		12: 2,
	}

	t := time.Now()
	hour, minutes, _ := t.Clock()
	roundMinutes := float64(minutes/5) * 5
	roundHours := hour % 12

	wantedPins := []int{1, 24, hourPins[roundHours]}
	wantedPins = append(wantedPins, minutePins[int(roundMinutes)]...)

	fmt.Println("Turning on pins", wantedPins)
	if usePins {
		for _, pinNumber := range wantedPins {
			pin := rpio.Pin(pinNumber)
			pin.Output()
			pin.High()
		}
	}

	allPins := []int{}
	for _, pins := range minutePins {
		allPins = append(allPins, pins...)
	}
	for _, pin := range hourPins {
		allPins = append(allPins, pin)
	}
	slices.Sort(allPins)
	allPins = slices.Compact(allPins)

	unwantedPins := slices.DeleteFunc(allPins, func(n int) bool {
		return slices.Contains(wantedPins, n)
	})

	fmt.Println("Turning off pins", unwantedPins)
	if usePins {
		for _, pinNumber := range unwantedPins {
			pin := rpio.Pin(pinNumber)
			pin.Low()
		}
	}

	if usePins {
		defer rpio.Close()
	}
}
