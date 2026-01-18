#include <stdio.h>
#include <time.h>

int main() {
    long long rounds = 3000000000LL;
     
    double sum = 0.0;
    double flip = -1.0;
    double pi;
    struct timespec tStart, tStop;
    long long ix;
    
    // Prime the caches.
    for (ix = 1; ix <= 3; ix++) {
        flip *= -1.0;
        sum += flip / (double)(ix + ix - 1);
    }
    sum = 0.0;
    flip = -1.0;

    // Timed test.
    clock_gettime(CLOCK_MONOTONIC, &tStart);    
    for (ix = 1; ix <= rounds; ix++) {
        flip *= -1.0;
        sum += flip / (double)(ix + ix - 1);
    }
    pi = sum * 4.0;
    clock_gettime(CLOCK_MONOTONIC, &tStop);
    
    // Report.
    double tDeltaSecs = (tStop.tv_sec - tStart.tv_sec) + 
                        (tStop.tv_nsec - tStart.tv_nsec) / 1e9;    
    
    printf("C,%lld,%.3f,%.16f\n", rounds, tDeltaSecs, pi); 
    
    return 0;
}
