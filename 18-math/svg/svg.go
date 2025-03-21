package svg

import (
	"fmt"
	"io"
	"time"

	clockface "github.com/thrashwerk/learn-go-with-tests-code/18-math"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
	 height="100%"
	 viewBox="0 0 300 300"
	 version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

// Write writes an SVG representation of an analogue clock, showing the time t, to the writer w.
func Write(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.SecondHandPoint(t), secondHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.MinuteHandPoint(t), minuteHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#0f0;stroke-width:5px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.HourHandPoint(t), hourHandLength)

	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#00f;stroke-width:7px;"/>`, p.X, p.Y)
}

func makeHand(p clockface.Point, length float64) clockface.Point {
	p = clockface.Point{X: p.X * length, Y: p.Y * length}             // scale
	p = clockface.Point{X: p.X, Y: -p.Y}                              // flip
	p = clockface.Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY} // translate

	return p
}
