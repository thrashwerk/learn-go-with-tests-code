package main

import (
	"os"
	"time"

	clockface "github.com/thrashwerk/learn-go-with-tests-code/18-math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
