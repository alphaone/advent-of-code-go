# Findings

In part 2:

- Only adding obstacles to positions the were on the original path reduced the
  runtime from ≈30secs to 7secs
- using `[][]rune` instead of `[]string` as `Grid` reduced runtime to 2.1secs
- using `*Grid` as method receiver did not do much
- using _go routines_ when trying to check loop candidates reduced the runtime
  to 0.4secs
