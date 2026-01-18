public class main {
    public static void main(String[] args) {
        long rounds = 3_000_000_000L;
        System.out.printf("Number of rounds: %d\n", rounds);

        double sum = 0.0;
        double flip = -1.0;
        long tStart, tStop;
        double pi;
        long ix;
        
        // Prime the caches.
        for (ix = 1; ix <= rounds; ix++) {
            flip *= -1.0;
            sum += flip / (double) (ix + ix - 1);
        }
        sum = 0.0;
        flip = -1.0;

        // Timed test.
        tStart = System.nanoTime();
        for (ix = 1; ix <= rounds; ix++) {
            flip *= -1.0;
            sum += flip / (double) (ix + ix - 1);
        }
        pi = sum * 4.0;
        tStop = System.nanoTime();

        // Report.
        double tDeltaSecs = (tStop - tStart) / 1e9;
        System.out.printf("Elapsed time (s): %.3f%n", tDeltaSecs);
        System.out.printf("Pi observed: %.16f\n", pi);
        System.out.printf("Pi expected: 3.1415926535897932\n");
    }
}

