use std::time::Instant;
use std::env;

fn main() {

    // Get the number of rounds.
    let str_rounds = env::var("rounds").unwrap_or_else(|_| {
        eprintln!("*** The 'rounds' environment variable is not set");
        std::process::exit(1);
    });
    let rounds: i64 = match str_rounds.parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("*** The 'rounds' environment variable must be an integer, saw: {}", str_rounds);
            std::process::exit(1);
        }
    };
   
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
    println!("Rust,{},{:.3},{:.40}", rounds, elapsed_secs, pi);
}
