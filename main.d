import std.stdio;
import std.process;
import std.conv;
import std.datetime.stopwatch;
import core.stdc.stdlib;

void main() {

    // Get the number of rounds.
    string strRounds = environment.get("rounds");
    if (strRounds is null) {
        writeln("*** The 'rounds' environment variable is not set");
        exit(1); 
    }
    long rounds;
    try {
        rounds = to!long(strRounds);
    } catch (ConvException e) {
        writefln("*** The 'rounds' environment variable must be an integer, saw: %s", strRounds);
        exit(1); 
    }
    
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
    writefln("D,%d,%.3f,%.40f", rounds, tDeltaSecs, pi);

}
