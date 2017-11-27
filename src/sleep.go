package main

import "time"
import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(seconds int) {
			time.Sleep(time.Duration(seconds) * time.Second)
			fmt.Println("Slept for:", seconds, "seconds")
			wg.Done()
		}(i)
	}
	fmt.Println("Finished in main.")
	wg.Wait()
}
