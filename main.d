import std.stdio;
import std.datetime.stopwatch;

void main() {
    long rounds = 3_000_000_000;
    writefln("Number of rounds: %d", rounds);
    
    double sum = 0.0;
    double flip = -1.0;
    double pi;
    long ix;
    
    // Prime the caches.
    for (ix = 1; ix <= 3; ix++) {
        flip *= -1.0;
        sum += flip / cast(double)(ix + ix - 1);
    }
    sum = 0.0;
    flip = -1.0;
    
    // Timed test.
    auto sw = StopWatch(AutoStart.yes);
    
    for (ix = 1; ix <= rounds; ix++) {
        flip *= -1.0;
        sum += flip / cast(double)(ix + ix - 1);
    }
    
    pi = sum * 4.0;
    sw.stop();
    
    // Report.
    double tDeltaSecs = sw.peek().total!"nsecs" / 1e9;
    writefln("Elapsed time (s): %.3f", tDeltaSecs);
    writefln("Pi observed: %.16f", pi);
    writefln("Pi expected: 3.1415926535897932");

}
