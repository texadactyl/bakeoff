import std/[times, os, strutils, strformat]

proc main() =
  # Get the number of rounds.
  let strRounds = getEnv("rounds")
  var rounds: int64
  
  try:
    rounds = parseBiggestInt(strRounds)
  except ValueError:
    echo "*** The 'rounds' environment variable must be an integer, saw: ", strRounds
    quit(1)
  
  var sum = 0.0
  var flip = -1.0
  var pi: float64
  var tStart: Time
  var tStop: Time
  
  # Prime the caches.
  for ix in 1'i64..3'i64:
    flip *= -1.0
    sum += flip / float64(ix + ix - 1)
  
  sum = 0.0
  flip = -1.0
  
  # Timed test.
  tStart = getTime()
  for ix in 1'i64..rounds:
    flip *= -1.0
    sum += flip / float64(ix + ix - 1)
  
  pi = sum * 4.0
  tStop = getTime()
  
  # Report.
  let tDeltaSecs = (tStop - tStart).inMilliseconds.float / 1000.0
  echo fmt"Nim,{rounds},{tDeltaSecs:.3f},{pi:.40f}"

when isMainModule:
  main()
