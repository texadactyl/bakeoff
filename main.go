package main

import (
    "time"
    "fmt"
)

func main() {
	rounds := int64(3_000_000_000)
    fmt.Printf("Number of rounds: %d\n", rounds);

	sum := 0.0
	flip := -1.0
	var pi float64;
	var tStart time.Time
	var tStop time.Time
	var ix int64

    // Prime the caches.
    for ix = int64(1); ix <= 3; ix++ {
        flip *= -1.0
        sum += flip / float64((ix + ix - 1))
    }
    sum = 0.0
    flip = -1.0
	
	// Timed test.
	tStart = time.Now()
    for ix = int64(1); ix <= rounds; ix++ {
        flip *= -1.0
        sum += flip / float64((ix + ix - 1))
    }
	pi = sum * 4.0
    tStop = time.Now()
    
    // Report.
    tDeltaSecs := (float64) ((tStop.UnixMilli() - tStart.UnixMilli())) / 1000.0
    fmt.Printf("Elapsed time (s): %.3f\n", tDeltaSecs)
	fmt.Printf("%.16f\n", pi)
	fmt.Printf("Pi expected: 3.1415926535897932\n");

}
