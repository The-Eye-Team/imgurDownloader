package main

import (
	"os"
	"sync"

	"github.com/labstack/gommon/color"
)

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func main() {
	var worker sync.WaitGroup
	var count, index int

	// Parse arguments
	parseArgs(os.Args)

	for {
		worker.Add(1)
		count++
		//go downloadID(&worker)
		index++
		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}
}
