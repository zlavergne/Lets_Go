/*
  The `original` code example can be found at https://play.golang.org/p/9U22NfrXeq
  from the "Go Concurrency Patterns" talk from 2012 found at
  https://talks.golang.org/2012/concurrency.slide#1

  An incredible visualisation can be found at http://divan.github.io/demos/primesieve/
*/
package main

// Fill the write-only channel `out` with integers
func fill(out chan<- int) {
	for i := 2; ; i++ {
		out <- i
	}
}

// Copy values from the read-only channel in
// If the value is not divisible by prime
// Copy to the write-only channel out
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in      // Copy value from 'in'.
		if i % prime != 0 {
			out <- i   // Forward 'i' to 'out'.
		}
	}
}

func main() {
	n := 100                   // The first `n` primes that we want to calculate
	ch := make(chan int)       // Create a new channel (bidirectional).
	go fill(ch)                // Launch a goroutine that fills ch.
	for i := 0; i < n; i++ {
		prime := <-ch          // Copy the first value from ch and treat it as our new prime
		print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
