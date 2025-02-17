package main

import (
	"os"
	"time"

	"github.com/thrashwerk/learn-go-with-tests-code/18-math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
