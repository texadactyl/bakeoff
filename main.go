package main

import (
	"bakeoff/current"
	"bakeoff/slotty"
	"fmt"
	"os"
	"strconv"
	"time"
)

const flagSlotty = true
const flagCurrent = true

var slottyETsecs float64
var currentETsecs float64

// Show help and then exit to the O/S
func showHelp() {
	fmt.Printf("\nUsage:  bakeoff  NLOOPS\n")
	fmt.Printf("\twhere NLOOPS is the number of loops to perform\n")
	fmt.Print("\n")
	os.Exit(0)
}

func main() {

	numLoops := -1
	dupAndDup2Divisor := 10

	var Args []string
	Args = append(Args, os.Args[1:]...)

	// If no options specified, just show help.
	if len(Args) != 1 {
		showHelp()
	}
	if Args[0] == "-h" {
		showHelp()
	}

	// Get numLoops.
	numLoops, err := strconv.Atoi(Args[0])
	if err != nil {
		fmt.Printf("Invalid numLoops: %s\n", Args[0])
		showHelp()
	}
	if numLoops < 1 {
		fmt.Printf("numLoops must be greater than zero\n")
		showHelp()
	}

	// If running the current Jacobin model, do so now.
	if flagCurrent {
		fmt.Println("Current begin .....")
		ti := time.Now()
		success := current.Current(numLoops, dupAndDup2Divisor)
		tf := time.Now()
		elapsed := tf.Sub(ti)
		currentETsecs = float64(elapsed.Milliseconds()) / 1000.0
		if !success {
			fmt.Printf("*** Current failed\n")
			os.Exit(86)
		}
		fmt.Printf("Current success, elapsed time: %.1f sec\n", currentETsecs)
	}

	// If running the parameter-slot-oriented Jacobin model, do so now.
	if flagSlotty {
		fmt.Println("Slotty begin .....")
		ti := time.Now()
		success := slotty.Slotty(numLoops, dupAndDup2Divisor)
		tf := time.Now()
		elapsed := tf.Sub(ti)
		slottyETsecs = float64(elapsed.Milliseconds()) / 1000.0
		if !success {
			fmt.Printf("*** Slotty failed\n")
			os.Exit(86)
		}
		fmt.Printf("Slotty success, elapsed time: %.1f sec\n", slottyETsecs)
	}

	// Show elapsed time savings if both ran.
	if flagCurrent && flagSlotty {
		diff := currentETsecs - slottyETsecs
		pctSaving := (diff / currentETsecs) * 100.0
		fmt.Printf("Slotty saves %.1f percent from the current methodology\n", pctSaving)
	}

}
