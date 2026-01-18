use std::time::Instant;

fn main() {
    let rounds: i64 = 3_000_000_000;
    println!("Number of rounds: {}", rounds);
    
    let mut sum: f64;  // ← No = 0.0 here
    let mut flip: f64 = -1.0;
    let pi: f64;
    
    // Prime the caches.
    sum = 0.0;  // ← First assignment here
    for ix in 1..=3 {
        flip *= -1.0;
        sum += flip / (ix + ix - 1) as f64;
    }
    let _watermelon = sum;
    let _cantaloupe = flip;
    sum = 0.0;
    flip = -1.0;
    
    // Timed test.
    let start = Instant::now();
    
    for ix in 1..=rounds {
        flip *= -1.0;
        sum += flip / (ix + ix - 1) as f64;
    }
    
    pi = sum * 4.0;
    let elapsed = start.elapsed();
    
    // Report.
    let elapsed_secs = elapsed.as_secs_f64();
    println!("Elapsed time (s): {:.3}", elapsed_secs);
    println!("Pi observed: {:.16}", pi);
    println!("Pi expected: 3.1415926535897932");
}
