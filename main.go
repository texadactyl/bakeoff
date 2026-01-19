package main

import (
    "time"
    "fmt"
    "os"
    "strconv"
)

func main() {

    // Get the number of rounds.
    strRounds := os.Getenv("rounds")
    rounds, err := strconv.ParseInt(strRounds, 10, 64)
        if err != nil {
            fmt.Printf("*** The 'rounds' environment variable must be an integer, saw: %s\n", strRounds)
            os.Exit(1)
        }
     
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
	fmt.Printf("Go,%d,%.3f,%.40f\n", rounds, tDeltaSecs, pi); 

}
