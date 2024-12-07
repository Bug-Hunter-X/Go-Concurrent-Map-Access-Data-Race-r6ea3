# Go Concurrent Map Access Data Race

This example demonstrates a common data race in Go when accessing and modifying a map from multiple goroutines concurrently without proper synchronization.  The program attempts to increment a map value 1000 times, but due to race conditions, the final result is often less than 1000.

## How to reproduce

1.  Save the code in `bug.go`
2.  Run `go run bug.go`
3.  Observe that the output is not always 1000, showcasing the data race.

## Solution

The solution involves using a `sync.Mutex` to protect access to the map.