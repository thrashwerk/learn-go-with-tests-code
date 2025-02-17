package clockface

import (
	"math"
	"time"
)

const (
	secondsInClock     = 60
	secondsInHalfClock = secondsInClock / 2
	minutesInClock     = 60
	minutesInHalfClock = minutesInClock / 2
	hoursInClock       = 12
	hoursInHalfClock   = hoursInClock / 2
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (math.Pi / (minutesInHalfClock / float64(t.Minute()))) + (secondsInRadians(t) / minutesInClock)
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock))) + (minutesInRadians(t) / hoursInClock)
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
