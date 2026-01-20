import time
import os
import sys

def main():
    # Get the number of rounds.
    str_rounds = os.getenv("rounds")

    try:
        rounds = int(str_rounds)
    except (ValueError, TypeError):
        print(f"*** The 'rounds' environment variable must be an integer, saw: {str_rounds}")
        sys.exit(1)

    sum_val = 0.0
    flip = -1.0

    # Prime the caches.
    for ix in range(1, 4):
        flip *= -1.0
        sum_val += flip / float(ix + ix - 1)

    sum_val = 0.0
    flip = -1.0

    # Timed test.
    t_start = time.time()
    for ix in range(1, rounds + 1):
        flip *= -1.0
        sum_val += flip / float(ix + ix - 1)

    pi = sum_val * 4.0
    t_stop = time.time()

    # Report.
    t_delta_secs = t_stop - t_start
    print(f"Python,{rounds},{t_delta_secs:.3f},{pi:.40f}")

if __name__ == "__main__":
    main()
